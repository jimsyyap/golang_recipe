package pkgquery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestPackageServer() *httptest.Server {
	pkgData := `[
        {"name": "package1", "version": "1.1"},        
        {"name": "package2", "version": "1.0"}        
    ]`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, pkgData)
	}))
	return ts
}

func TestFetchPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()

	packages, err := fetchPackageData(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(packages) != 2 {
		t.Errorf("Expected 2 packages, got %d", len(packages))
	}
}
