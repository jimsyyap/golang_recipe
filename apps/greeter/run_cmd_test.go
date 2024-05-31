package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		output string
		err    error
	}{
		{
			c:      config{printUsage: true},
			output: usageString,
		},
		{
			c:      config{numTimes: 5},
			input:  "",
			output: strings.Repeat("name pls\n", 5),
			err:    errors.New("enter name pls"),
		},
		{
			c:      config{numTimes: 5},
			input:  "Bill Bryson\n",
			output: "Your name pls, enter when done\n" + strings.Repeat("nice to meet you", 5),
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		r := strings.NewReader(tc.input)
		err := runCmd(r, byteBuf, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if tc.err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("Expected error %v, got %v", tc.err, err)
			}
		}
		gotMsg := byteBuf.String()
		if gotMsg != tc.output {
			t.Errorf("Expected output %v, got %v", tc.output, gotMsg)
		}
	}
}

/*
* this code makes sure that Cmd plays each of the cases in the test correctly
* it's like a checklist for the code before it goes out to production.
 */
