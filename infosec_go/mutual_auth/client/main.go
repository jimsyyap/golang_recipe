package main

import (
    "crypto/tls"
    "cripto/x509"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    var (
        err error
        cert tls.Certificate
        serverCert, body []byte
        pool *x509.CertPool
        tlsConf *tls.Config
        transport *http.transport
        client *http.client
        resp *http.Response
    )

    if cert, err = tls.LoadX509KeyPair("cert.pem", "key.pem"); err != nil {
        log.Fatal(err)
    }

    if serverCert, err = ioutil.ReadFile("server.crt"); err != nil {
        log.Fatal(err)
    }

    pool = x509.NewCertPool()
    pool.AppendCertsFromPEM(serverCert)

    tlsConf = &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs: pool,
    }

    transport = &http.Transport{
        TLSClientConfig: tlsConf,
    }

    client = &http.Client{Transport: transport}

    if resp, err = client.Get(https://localhost:8080); err != nil {
        log.Fatal(err)
    }
    if body, err = ioutil.ReadAll(resp.Body); err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    fmt.Printf("%s\n", body)
}
