package html

import (
	"github.com/jmizell/BranchDecoder/decoder/encodings"
	"github.com/jmizell/GoLibraryRewrites/html"
)

var Decoder *decoder

type decoder struct {
}

func (d *decoder) Type() encodings.Encoding {
	return encodings.HTML
}

func (d *decoder) Decode(data []byte) (bool, []byte) {
	return html.UnescapeByte(data)
}