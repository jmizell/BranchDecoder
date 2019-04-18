package url

import (
	"testing"
)

func TestDecoder_Detect(t *testing.T) {

	t.Run("plus_encode", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("one+two")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("percent_encode", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("one%20two")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("no_encoding", func(subTest *testing.T) {
		if Decoder.Detect([]byte("one")) {
			subTest.Fatalf("expected no match")
		}
	})

	t.Run("invalid_character", func(subTest *testing.T) {
		if Decoder.Detect([]byte("one%20two=")) {
			subTest.Fatalf("expected no match")
		}
	})

	t.Run("space", func(subTest *testing.T) {
		if Decoder.Detect([]byte("one%20 two")) {
			subTest.Fatalf("expected no match")
		}
	})
}

func TestDecoder_Decode(t *testing.T) {

	t.Run("no_encoding", func(subTest *testing.T) {
		if m, _ := Decoder.Decode([]byte("test")); m {
			subTest.Fatalf("expected no match")
		}
	})

	t.Run("plus_encoded", func(subTest *testing.T) {
		if m, value := Decoder.Decode([]byte("one+two")); !m {
			subTest.Fatalf("expected match")
		} else if string(value) != "one two" {
			subTest.Fatalf("expected \"one two\", received \"%s\"", string(value))
		}
	})

	t.Run("percent_encoded", func(subTest *testing.T) {
		if m, value := Decoder.Decode([]byte("one%20two")); !m {
			subTest.Fatalf("expected match")
		} else if string(value) != "one two" {
			subTest.Fatalf("expected \"one two\", received \"%s\"", string(value))
		}
	})

	t.Run("invalid_encoding", func(subTest *testing.T) {
		if m, _ := Decoder.Decode([]byte("one%20=two")); m {
			subTest.Fatalf("expected no match")
		}
	})
}

func BenchmarkDetectURLEncoding(b *testing.B) {
	examples := [][]byte{
		[]byte(`Testing%20One%20%2B%20Two%20three%3F%3F%20Test%21`),
		[]byte(`Testing=One%20%2B%20Two%20three%3F%3F%20Test%21`),
	}

	for i := 0; i < b.N; i++ {
		for _, e := range examples {
			Decoder.Detect(e)
		}
	}
}

func BenchmarkDecodeURLEncoding(b *testing.B) {
	examples := [][]byte{
		[]byte(`Testing%20One%20%2B%20Two%20three%3F%3F%20Test%21`),
		[]byte(`Testing=One%20%2B%20Two%20three%3F%3F%20Test%21`),
	}

	for i := 0; i < b.N; i++ {
		for _, e := range examples {
			Decoder.Decode(e)
		}
	}
}
