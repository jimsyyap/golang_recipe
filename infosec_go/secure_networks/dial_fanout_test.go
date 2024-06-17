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
}
