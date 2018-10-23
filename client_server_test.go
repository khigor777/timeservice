package timeservice

import (
	"context"
	"testing"
	"time"
)

func TestClientServer(t *testing.T) {
	ts := NewServer("8080")
	go func() {
		time.Sleep(time.Second)
		res, err := TcpClientRun("127.0.0.1", "8080")

		if err != nil {
			t.Errorf("error in tcp client: %s", err.Error())
		}

		if res == 0 {
			t.Errorf("empty response ")
		}

	}()

	go ts.Run()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	time.Sleep(time.Second * 2)
	ts.Shutdown(ctx)
}
