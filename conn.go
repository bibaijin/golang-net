package main

import (
	"bufio"
	"io"
	"net"
	"sync"
	"time"

	"github.com/golang/glog"
)

type TimeoutConn struct {
	timeout time.Time
	conn    net.Conn
}

func NewTimeoutConn(timeout time.Time, conn net.Conn) *TimeoutConn {
	tc := &TimeoutConn{
		timeout: timeout,
		conn:    conn,
	}
	return tc
}

func (tc *TimeoutConn) Read(b []byte) (n int, err error) {
	tc.conn.SetReadDeadline(tc.timeout)
	n, err = tc.conn.Read(b)
	return n, err
}

func (tc *TimeoutConn) Write(b []byte) (n int, err error) {
	tc.conn.SetWriteDeadline(tc.timeout)
	n, err = tc.conn.Write(b)
	return n, err
}

func (tc *TimeoutConn) Close() error {
	return tc.conn.Close()
}

func (tc *TimeoutConn) LocalAddr() net.Addr {
	return tc.conn.LocalAddr()
}

func (tc *TimeoutConn) RemoteAddr() net.Addr {
	return tc.conn.RemoteAddr()
}

// func (tc *timeoutConn) ReadString
