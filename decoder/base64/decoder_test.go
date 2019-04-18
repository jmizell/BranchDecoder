package base64

import (
	"testing"
)

func TestDecoder_Detect(t *testing.T) {

	t.Run("standard_padded", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY+KJheKKl868xZLigKHimaA=")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("standard_unpadded", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY+KJheKKl868xZLigKHimaA")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("url_padded", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY-KJheKKl868xZLigKHimaA=")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("url_standard_invalid_mix", func(subTest *testing.T) {
		if Decoder.Detect([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY+-KJheKKl868xZLigKHimaA=")) {
			subTest.Fatalf("expected no match")
		}
	})

	t.Run("url_unpadded", func(subTest *testing.T) {
		if !Decoder.Detect([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY-KJheKKl868xZLigKHimaA")) {
			subTest.Fatalf("expected match")
		}
	})

	t.Run("invalid", func(subTest *testing.T) {
		if Decoder.Detect([]byte("!dGVzdA==")) {
			subTest.Fatalf("expected no match")
		}
	})
}

func TestDecoder_Decode(t *testing.T) {

	expectedValue := "abc≅⊗μŒ‡♠abc≅⊗μŒ‡♠"

	t.Run("standard_padded", func(subTest *testing.T) {
		m, value := Decoder.Decode([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY+KJheKKl868xZLigKHimaA="))
		if !m {
			subTest.Fatalf("expected match")
		} else if string(value) != expectedValue {
			subTest.Fatalf("expected to decode value to \"%s\", but found \"%s\"", expectedValue, string(value))
		}
	})

	t.Run("standard_unpadded", func(subTest *testing.T) {
		m, value := Decoder.Decode([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY+KJheKKl868xZLigKHimaA"))
		if !m {
			subTest.Fatalf("expected match")
		} else if string(value) != expectedValue {
			subTest.Fatalf("expected to decode value to \"%s\", but found \"%s\"", expectedValue, string(value))
		}
	})

	t.Run("url_padded", func(subTest *testing.T) {
		m, value := Decoder.Decode([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY-KJheKKl868xZLigKHimaA="))
		if !m {
			subTest.Fatalf("expected match")
		} else if string(value) != expectedValue {
			subTest.Fatalf("expected to decode value to \"%s\", but found \"%s\"", expectedValue, string(value))
		}
	})

	t.Run("url_unpadded", func(subTest *testing.T) {
		m, value := Decoder.Decode([]byte("YWJj4omF4oqXzrzFkuKAoeKZoGFiY-KJheKKl868xZLigKHimaA"))
		if !m {
			subTest.Fatalf("expected match")
		} else if string(value) != expectedValue {
			subTest.Fatalf("expected to decode value to \"%s\", but found \"%s\"", expectedValue, string(value))
		}
	})

	t.Run("invalid", func(subTest *testing.T) {
		m, _ := Decoder.Decode([]byte("!dGVzdA=="))
		if m {
			subTest.Fatalf("expected no match")
		}
	})
}
