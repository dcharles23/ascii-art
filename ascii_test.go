package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPrintAllStringASCII(t *testing.T) {
	testCases := []struct {
		inputFile          string
		expectedOutputFile string
	}{
		{"testdata/input1.txt", "testdata/expected_output1.txt"},
		{"testdata/input2.txt", "testdata/expected_output2.txt"},
		{"testdata/input3.txt", "testdata/expected_output3.txt"},
		{"testdata/input4.txt", "testdata/expected_output4.txt"},
		{"testdata/input5.txt", "testdata/expected_output5.txt"},
		{"testdata/input6.txt", "testdata/expected_output6.txt"},
		{"testdata/input7.txt", "testdata/expected_output7.txt"},
		{"testdata/input8.txt", "testdata/expected_output8.txt"},
		{"testdata/input9.txt", "testdata/expected_output9.txt"},
		{"testdata/input10.txt", "testdata/expected_output10.txt"},
		{"testdata/input11.txt", "testdata/expected_output11.txt"},
		{"testdata/input12.txt", "testdata/expected_output12.txt"},
		{"testdata/input13.txt", "testdata/expected_output13.txt"},
		{"testdata/input14.txt", "testdata/expected_output14.txt"},
		{"testdata/input15.txt", "testdata/expected_output15.txt"},
		{"testdata/input16.txt", "testdata/expected_output16.txt"},
		{"testdata/input17.txt", "testdata/expected_output17.txt"},
	}

	for _, tc := range testCases {
		// Load input and expected output files
		input, _ := ioutil.ReadFile(tc.inputFile)
		expectedOutput, _ := ioutil.ReadFile(tc.expectedOutputFile)

		fileLines := ReadStandardTxt()
		asciiTemplates := return2dASCIIArray(fileLines)

		// Call the function
		output := capturePrint(func() {
			printAllStringASCII(string(input), asciiTemplates)
		})

		// Compare the actual and expected output
		if strings.TrimSpace(output) != strings.TrimSpace(string(expectedOutput)) {
			t.Errorf("Expected output does not match actual output for input %s", tc.inputFile)
		}
	}
}

// Utility function to capture the printed output
func capturePrint(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old
	out, _ := ioutil.ReadAll(r)
	return string(out)
}