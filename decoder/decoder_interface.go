package decoder

import (
	"github.com/jmizell/BranchDecoder/decoder/encodings"
)

// Decoder is the interface a Decoder must implement
// to be used in BranchDecoder
type Decoder interface {
	Detect([]byte) bool
	Decode([]byte) (bool, []byte)
	Type() encodings.Encoding
}
