package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	// Not closing the listener can lead to memory leaks or deadlocks
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}
