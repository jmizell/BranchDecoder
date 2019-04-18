package html

import "testing"

func TestDecoder_Detect(t *testing.T) {

	t.Run("hex", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("&#x00026;")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("number", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("&#38;")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("name", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("&amp;")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("double_encoded", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("&amp;#10026;")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("no_match", func(subTest *testing.T) {
		if Decoder.Detect([]byte("&;")) {
			subTest.Fatalf("expected no match")
		}
	})
}

func TestDecoder_Decode(t *testing.T) {

	t.Run("hex", func(subTest *testing.T) {
		if m, value := Decoder.Decode([]byte("&#x00026;")); !m {
			subTest.Fatalf("expected match")
		} else if string(value) != "&" {
			subTest.Fatalf("expected to decode to &, found %s", string(value))
		}
	})

	t.Run("number", func(subTest *testing.T) {
		if m, value := Decoder.Decode([]byte("&#38;")); !m {
			subTest.Fatalf("expected match")
		} else if string(value) != "&" {
			subTest.Fatalf("expected to decode to &, found %s", string(value))
		}
	})

	t.Run("name", func(subTest *testing.T) {
		if m, value := Decoder.Decode([]byte("&amp;")); !m {
			subTest.Fatalf("expected match")
		} else if string(value) != "&" {
			subTest.Fatalf("expected to decode to &, found %s", string(value))
		}
	})

	t.Run("match_no_encoding", func(subTest *testing.T) {
		if m, value := Decoder.Decode([]byte("&;")); !m {
			subTest.Fatalf("expected match")
		} else if string(value) != "&;" {
			subTest.Fatalf("expected to decode to &;, found %s", string(value))
		}
	})

	t.Run("no_match", func(subTest *testing.T) {
		if m, _ := Decoder.Decode([]byte("test")); m {
			subTest.Fatalf("expected no match")
		}
	})
}
