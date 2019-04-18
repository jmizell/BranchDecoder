package url

import (
	"github.com/jmizell/BranchDecoder/decoder/encodings"
)

var Decoder *decoder

type decoder struct {
}

func (d *decoder) Type() encodings.Encoding {
	return encodings.URL
}
