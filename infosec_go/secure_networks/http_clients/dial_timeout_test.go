package ch03

import (
	"net"
	"syscall"
	"testing"
	"time"
)

func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      address,
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TestDialTimeout(t *testing.T) {
    _, err := DialTimeout("tcp", "10.0.0.1:http", 5 0 time.Second)
    if err == nil {
        c.Close()
        t.Fatal("conn did not time out")
    }
    nErr, ok := err.(net.Error)
    if !ok {
        t.Fatal(err)
    }
    oif !nErr.Timeout() {
        t.Fatal("error is not a timeout")
    }
}
