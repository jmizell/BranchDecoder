package decoder

import (
	"encoding/json"
	"github.com/jmizell/BranchDecoder/decoder/encodings"
)

// EncodingTree is the record of matched encodings, and
// intermediate decoded values, if enabled.
type EncodingTree struct {
	Encoding        encodings.Encoding
	DecodedValue    []byte
	NestedEncodings []*EncodingTree
}

type intermediateJSONEncodingTree struct {
	Encoding        encodings.Encoding `json:"encoding"`
	DecodedValue    string             `json:"decoded_value"`
	NestedEncodings []*EncodingTree    `json:"nested_encodings,omitempty"`
}

// MarshalJSON will output EncodingTree with DecodedValue as
// a string, instead of a byte array.
func (e *EncodingTree) MarshalJSON() ([]byte, error) {

	s := &intermediateJSONEncodingTree{
		Encoding:        e.Encoding,
		DecodedValue:    string(e.DecodedValue),
		NestedEncodings: e.NestedEncodings,
	}

	return json.MarshalIndent(s, "", "  ")
}

// UnmarshalJSON will decode a json EncodingTree with
// DecodedValue as a string to the byte array value.
func (e *EncodingTree) UnmarshalJSON(data []byte) error {

	s := &intermediateJSONEncodingTree{}
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	e.DecodedValue = []byte(s.DecodedValue)
	e.Encoding = s.Encoding
	e.NestedEncodings = s.NestedEncodings

	return nil
}


func (e *EncodingTree) String() string {

	d, _ := json.MarshalIndent(e, "", "  ")
	return string(d)
}