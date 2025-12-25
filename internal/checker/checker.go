package checker

import (
	"context"
	"fmt"
	"net"
)

const (
	networkProtocol = "tcp"
	statusClosed    = "closed"
	statusOpen      = "open"
)

// PortStatus represents the result of a port check.
type PortStatus struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Status string `json:"status"`
}

// TCPChecker implements Service using net.Dialer.
type TCPChecker struct {
	dialer *net.Dialer
}

// NewTCPChecker creates a new instance of TCPChecker.
func NewTCPChecker() *TCPChecker {
	return &TCPChecker{
		dialer: &net.Dialer{},
	}
}

func (c *TCPChecker) Check(ctx context.Context, ip string, port int) PortStatus {
	target := fmt.Sprintf("%s:%d", ip, port)

	// We use the context for timeout
	conn, err := c.dialer.DialContext(ctx, networkProtocol, target)

	status := statusClosed
	if err == nil {
		status = statusOpen
		_ = conn.Close()
	}

	return PortStatus{
		IP:     ip,
		Port:   port,
		Status: status,
	}
}
