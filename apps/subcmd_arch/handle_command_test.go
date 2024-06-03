// Listing 2.6: chap2/sub-cmd-arch/handle_command_test.go
package main

import (
	"bytes"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	usageMessage := `Usage: mync [http|grpc] -h
    http: A HTTP client.

    http: <options> server

    Options:
        -verb string
            HTTP method (default "GET")
    grpc: A gRPC client.
    Grpc: <options> server

    Options:
        -body string
            Body of request
        -method string
            Method of request
    `
	testConfigs := []struct {
		args   []string
		output string
		err    error
	}{
		{
			args:   []string{},
			err:    errInvalidSubCommand,
			output: "invalid sub-command specified\n" + usageMessage,
		},
		{
			args:   []string{"-h"},
			err:    nil,
			output: usageMessage,
		},
		{
			args:   []string{"foo"},
			err:    errInvalidSubCommand,
			output: "invalid sub-command specified\n" + usageMessage,
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range testConfigs {
		err := handleCommand(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}

		if tc.err != nil && err == nil {
			t.Fatalf("Expected error %v, got nil", tc.err)
		}

		if len(tc.output) != 0 {
			gotOutput := byteBuf.String()
			if gotOutput != tc.output {
				t.Fatalf("Expected output %q, got %q", tc.output, gotOutput)
			}
		}
		byteBuf.Reset()
	}
}
