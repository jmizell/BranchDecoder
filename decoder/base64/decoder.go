package base64

import (
	"bytes"
	"encoding/base64"

	"github.com/jmizell/BranchDecoder/decoder/encodings"
)

var Decoder *decoder

type decoder struct {
}

func (d *decoder) Type() encodings.Encoding {
	return encodings.URL
}

// Decode decodes base64 data in formats specified in RFC 4648
// TODO support variants https://en.wikipedia.org/wiki/Base64#Variants_summary_table
func (d *decoder) Decode(data []byte) (bool, []byte) {

	// select the right encoder for the input
	padded := bytes.HasSuffix(data, []byte("="))
	urlEncoding := bytes.ContainsAny(data, "-_")
	encoder := base64.RawStdEncoding
	switch {
	case padded && urlEncoding:
		encoder = base64.URLEncoding
	case padded && !urlEncoding:
		encoder = base64.StdEncoding
	case !padded && urlEncoding:
		encoder = base64.RawURLEncoding
	}

	// decode our data to our destination
	dst := make([]byte, encoder.DecodedLen(len(data)))
	_, err := encoder.Decode(dst, data)
	if err != nil {
		return false, nil
	}

	// trim excess space from the end of the array
	dst = bytes.TrimRightFunc(dst, func(r rune) bool {
		return r == 0
	})

	return true, dst
}
