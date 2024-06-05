package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, client")
	}))
	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	expected := "hello, client"
	data, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if expected != string(data) {
		t.Errorf("expected %s, got %s", expected, string(data))
	}

}
