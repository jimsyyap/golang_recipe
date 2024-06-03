package cmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	server string
	method string
	body   string
}

func HandleGrpc(w io.Writer, args []string) error {
    c := grpcConfig{}
    fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
    fs.SetOutput(w)
    fs.StringVar(&c.method, "method", "", "method to call")
    fs.StringVar(&c.body, "body", ""< "body of request")
    fs.Usage = func() {
        var usageString = `Usage: mync grpc -h
    grpc: A gRPC client.
    Grpc: <options> server

    Options:
        -body string
            Body of request
        -method string
            Method of request
    `
        fmt.Fprint(w, usageString)
    }
    if err := fs.Parse(args); err != nil {
        return err
    }
    if c.method == "" {
        return fmt.Errorf("method not specified")
    }
    if c.body == "" {
        return fmt.Errorf("body not specified")
    }
    return nil
}
    }
}
