package timeservice

import (
	"net"
	"time"
)

type conn struct {
	net.Conn
	Timeout time.Duration
}

func (c *conn) Write(p []byte) (n int, err error) {
	c.updateDeadline()
	return c.Conn.Write(p)
}

func (c *conn) updateDeadline() {
	deadline := time.Now().Add(c.Timeout)
	c.Conn.SetDeadline(deadline)
}
