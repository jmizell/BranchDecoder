package encodings

import (
	"encoding/json"
	"strings"
)

type Encoding uint64

const (
	Unknown Encoding = iota
	None
	URL
	HTML
)

func (e *Encoding) Parse(encoding string) {

	switch strings.ToLower(encoding) {
	case "url":
		*e = URL
	case "html":
		*e = HTML
	case "none":
		*e = None
	default:
		*e = Unknown
	}
}

func (e Encoding) String() string {

	switch e {
	case URL:
		return "URL"
	case HTML:
		return "HTML"
	case None:
		return "None"
	default:
		return "Unknown"
	}
}

func (e *Encoding) UnmarshalJSON(data []byte) error {

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	e.Parse(s)

	return nil
}

func (e *Encoding) MarshalJSON() ([]byte, error) {

	return json.Marshal(e.String())
}
