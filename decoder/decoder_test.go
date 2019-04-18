package decoder

import (
	"fmt"
	"testing"

	"github.com/jmizell/BranchDecoder/decoder/encodings"
)

var debugOutput = false

type decodeTestCase struct {
	Name          string
	UnencodedData string
	EncodedData   string
	ExpectedStack []encodings.Encoding
}

var decodeTestCases = []*decodeTestCase{

	/*
		No Encodings
	*/
	{
		Name:          "None[test,1]",
		UnencodedData: "test ✪", EncodedData: "test ✪",
		ExpectedStack: []encodings.Encoding{encodings.None},
	},

	/*
		Single URL Encoding
	*/
	{
		Name:          "URL[test,✪]",
		UnencodedData: "test ✪", EncodedData: "test%20%E2%9C%AA",
		ExpectedStack: []encodings.Encoding{encodings.URL},
	},

	/*
		Double URL Encoding
	*/
	{
		Name:          "URL,URL[test,✪]",
		UnencodedData: "test ✪", EncodedData: "test%2520%25E2%259C%25AA",
		ExpectedStack: []encodings.Encoding{encodings.URL, encodings.URL},
	},

	/*
		URL Encoding four times
	*/
	{
		Name:          "URL,URL,URL,URL[test,✪]",
		UnencodedData: "test ✪", EncodedData: "test%25252520%252525E2%2525259C%252525AA",
		ExpectedStack: []encodings.Encoding{encodings.URL, encodings.URL, encodings.URL, encodings.URL},
	},

	/*
		URL Encoding five times
	*/
	{
		Name:          "URL,URL,URL,URL,URL[test,✪]",
		UnencodedData: "test%20%E2%9C%AA", EncodedData: "test%2525252520%25252525E2%252525259C%25252525AA",
		ExpectedStack: []encodings.Encoding{encodings.URL, encodings.URL, encodings.URL, encodings.URL},
	},

	/*
		HTML Encoded
	*/
	{
		Name:          "HTML[test,✪]",
		UnencodedData: "test ✪", EncodedData: "test &#10026;",
		ExpectedStack: []encodings.Encoding{encodings.HTML},
	},
	{
		Name:          "HTML,HTML[test,✪]",
		UnencodedData: "test ✪", EncodedData: "test &amp;#10026;",
		ExpectedStack: []encodings.Encoding{encodings.HTML, encodings.HTML},
	},

	/*
		URL, HTML Encoded
	*/
	{
		Name:          "URL,HTML[test,✪]",
		UnencodedData: "test ✪", EncodedData: "test%20%26%2310026%3B",
		ExpectedStack: []encodings.Encoding{encodings.URL, encodings.HTML},
	},
}

func TestBranchDecoder_Decode(t *testing.T) {

	for _, testCase := range decodeTestCases {
		t.Run(testCase.Name, func(subTest *testing.T) {

			d := NewBranchDecoder()
			d.PreserveIntermediateValues = false
			d.MaxDepth = 3
			tree := d.Decode([]byte(testCase.EncodedData))

			if debugOutput {
				fmt.Println(tree.String())
			}

			if testCase.ExpectedStack[0] == encodings.None {

				if tree.Encoding != encodings.None {
					subTest.Fatalf("expected encoding None, found %s", tree.Encoding.String())
				}

				if string(tree.DecodedValue) != testCase.UnencodedData {
					subTest.Fatalf("expected=\"%s\" != decoded=\"%s\"", testCase.UnencodedData, string(tree.DecodedValue))
				}

				return
			}

			walkEncodingTree(subTest, tree, testCase, 0, 4)
		})
	}
}

func walkEncodingTree(t *testing.T, tree *EncodingTree, testCase *decodeTestCase, depth, maxDepth int) {

	// We're at the final layer, or maxdepth
	if depth > len(testCase.ExpectedStack) || depth == maxDepth {

		// if we're at the end of the encodings, it's None, if we've reached
		// max iteration depth, then we expect to find the last encoding
		expectedEncoding := encodings.None
		if depth == maxDepth {
			expectedEncoding = testCase.ExpectedStack[depth-1]
		}

		if tree.Encoding != expectedEncoding {
			t.Fatalf("expected to find final decoding of None at depth %d, found %s", depth, expectedEncoding)
		}

		if string(tree.DecodedValue) != testCase.UnencodedData {
			t.Fatalf("expected=\"%s\" != decoded=\"%s\"", testCase.UnencodedData, string(tree.DecodedValue))
		}

		return
	}

	// determine our expected encoding at this depth
	expectedEncoding := encodings.None
	if depth < len(testCase.ExpectedStack) {
		expectedEncoding = testCase.ExpectedStack[depth]
	}

	// search for our embedded encoding, and recurse
	for _, detectedEncoding := range tree.NestedEncodings {
		if detectedEncoding.Encoding == expectedEncoding {
			walkEncodingTree(t, detectedEncoding, testCase, depth+1, maxDepth)
			return
		}
	}

	t.Fatalf("failed to find %s encoding in stack at depth %d", expectedEncoding, depth)
}
