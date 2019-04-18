package decoder

import (
	"github.com/jmizell/BranchDecoder/decoder/encodings"
	"github.com/jmizell/BranchDecoder/decoder/html"
	"github.com/jmizell/BranchDecoder/decoder/url"
)

var decoders = map[encodings.Encoding]Decoder{
	encodings.URL: url.Decoder,
	encodings.HTML: html.Decoder,
}

// BranchDecoder is a recursive, broadly searching, string encoding detection,
// and decoding tool. It attempts to walk through every potential string
// encoding, and assembles a EncodingTree of matched encodings and optionally
// their intermediate values.
//
// BranchDecoder is intended to help detect encodings in ambiguous inputs,
// such as obfuscated payloads in form data.
//
// By default, BranchDecoder runs detection on each encoding at once; you can
// tune this setting by adjusting the Workers value to a lower number then the
// number of supported decoders. MaxDepth can limit the recursive depth
// BranchDecoder attempts to detect encodings.
type BranchDecoder struct {

	// MaxDepth is the maximum recursion depth
	// that BranchDecoder will search for encodings
	MaxDepth                   int

	// Workers is the number of decoding workers
	// to launch when Decode is called
	Workers                    int

	// PreserveIntermediateValues enables storing
	// intermediate decoded values
	PreserveIntermediateValues bool

	// Encodings is the Decoders supported by BranchDecoder,
	// if this value is unset, then the default decoder list
	// is used. If you want to limit the number of decoders
	// that BranchDecoder uses, you can do that by setting
	// this value.
	Encodings                  map[encodings.Encoding]Decoder
}

// NewBranchDecoder returns a new BranchDecoder with default values
func NewBranchDecoder() *BranchDecoder {
	return &BranchDecoder{
		MaxDepth:                   10,
		PreserveIntermediateValues: true,
	}
}

// Decode recursively searches through each possible encoding, and returns a EncodingTree
// of detected encodings.
func (b *BranchDecoder) Decode(data []byte) *EncodingTree {
	return b.decode(data, 0)
}

func (b *BranchDecoder) decode(data []byte, depth int) *EncodingTree {

	targetEncodings := b.Encodings
	if targetEncodings == nil {
		targetEncodings = decoders
	}

	workerCount := b.Workers
	if workerCount == 0 || workerCount > len(targetEncodings) {
		workerCount = len(targetEncodings)
	}

	jobs := make(chan encodings.Encoding, len(targetEncodings))
	results := make(chan *EncodingTree, len(targetEncodings))
	for w := 0; w < workerCount; w++ {
		go func() {
			for j := range jobs {

				// skip anything that cannot be decoded
				if !targetEncodings[j].Detect(data) {
					results <- nil
					continue
				}

				// attempt to decode value
				ok, value := targetEncodings[j].Decode(data)
				if !ok {
					results <- nil
					continue
				}

				// recurse decode
				var nestedEncodings []*EncodingTree
				if depth < b.MaxDepth {
					r := b.decode(value, depth+1)
					if r.Encoding == encodings.None {
						nestedEncodings = []*EncodingTree{r}
					} else {
						nestedEncodings = r.NestedEncodings
					}
				}

				et := &EncodingTree{
					Encoding:        j,
					NestedEncodings: nestedEncodings,
				}

				if b.PreserveIntermediateValues || depth >= b.MaxDepth {
					et.DecodedValue = value
				}

				results <- et
			}
		}()
	}

	for e := range targetEncodings {
		jobs <- e
	}
	close(jobs)

	tree := &EncodingTree{
		Encoding: encodings.Unknown,
	}
	for range targetEncodings {
		r := <-results
		if r != nil {
			tree.NestedEncodings = append(tree.NestedEncodings, r)
		}
	}

	if len(tree.NestedEncodings) == 0 {
		tree.Encoding = encodings.None
		tree.DecodedValue = data
	} else if b.PreserveIntermediateValues {
		tree.DecodedValue = data
	}

	return tree
}
