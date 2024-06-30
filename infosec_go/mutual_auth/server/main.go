package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("hello %s\n", r.TLS.PeerCertificates[0].Subject.CommonName)
    fmt.Fprint(w, "authenticated\n")
}

func main() {
    var (
        err error
        clientCert []byte
        pool *x509.CertPool
        tlsConfig *tls.Config
        server *http.Server
    )

    http.HandleFunc("/hello", helloHandler)

    if clientCert, err = ioutil.ReadFile("client.pem"); err != nil {
        log.Fatal(err)
    }

    pool = x509.NewCertPool()
    pool.AppendCertsFromPEM(clientCert)

    tlsConfig = &tls.Config{
        ClientCAs: pool,
        ClientAuth: tls.RequireAndVerifyClientCert,
    }
    tls.Config.BuildNameToCertificate()

    server = &http.Server{
        Addr: ":9443",
        TLSConfig: tlsConfig,
    }
    log.Fatal(server.ListenAndServeTLS("server.pem", "server-key.pem"))
}
