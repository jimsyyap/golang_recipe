package ch03

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"
)

func TestDialContextCancelFanout(t *testing.T) {
	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(10*time.Second),
	)

    listener, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        t.Fatal(err)
    }
    defer listener.Close()

    go func() {
        conn, err := listener.Accept()
        if err != nil {
            conn.Close()
        }
    }()

    dial := func(ctx context.Context, address string, response chan int, id int, wg *sync.Waitgroup) {
        defer wg.Done()

        var d net.Dialerc, err := d.DialContext(ctx, "tcp", address)
        if err != nil {
            return
        }
        c.Close()

        select {
        case <-ctx.Done():
        case response <- id:
        }
    }

    res := make(chan int)
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go dial(ctx, listener.Addr().String(), res, i, &wg)
    }

    response := <-res
    cancel()
    wg.Wait()
    close(res)

    if ctx.Err() != context.Canceled {
        t.Errorf("expected context to be canceled, got %s", ctx.Err())
    }

    t.Logf("dialer %d succeeded", response)
}
