package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

type SockConf struct {
	Timeout time.Duration // Timeout
	Buffer  int           // Buffer chunk, default 2048
}

type refConn struct {
	Write           func([]byte) (int, error)
	Read            func([]byte) (int, error)
	SetReadDeadline func(time.Time) error
	Close           func() error
}

type linkc struct {
	ref               refConn
	timeout           time.Duration
	host              string
	errormsg          error
	buffer            int
	file              string
	ContentLength     chan int64
	ContentDownloaded chan int64
}

func Link(url string, config SockConf) (linknode linkc) {
	addr, scheme, err := toValidAddr(url)
	if err != nil {
		panic(err)
	}

	if config.Buffer < 1 {
		config.Buffer = 2048
	}
	linkconfig := linkc{host: addr, buffer: config.Buffer, timeout: config.Timeout}
	
	if scheme == "https" {
		var ref *tls.Conn
		ref, err = tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: false})
		if err != nil {
			panic(errWarn(SOCKET_ERR, fmt.Sprintf("%s, %s", addr, err.Error())))
		}
		linkconfig.ref = refConn{Write: ref.Write, Read: ref.Read, SetReadDeadline: ref.SetReadDeadline, Close: ref.Close}
	} else {
		dialer := &net.Dialer{Timeout: config.Timeout}
		var ref net.Conn
		ref, err = dialer.Dial("tcp", addr)
		if err != nil {
			panic(errWarn(SOCKET_ERR, fmt.Sprintf("%s, %s", addr, err.Error())))
		}
		linkconfig.ref = refConn{Write: ref.Write, Read: ref.Read, SetReadDeadline: ref.SetReadDeadline, Close: ref.Close}
	}

	if err == nil {
		linkconfig.ContentLength = make(chan int64)
		linkconfig.ContentDownloaded = make(chan int64)

		return linkconfig
	}
	panic(errWarn(SOCKET_ERR, err.Error()))
}

func (conn *linkc) Req(method reqmethod, body BodyRequest) (success bool, response BodyResponse) {
	var withPayload bool
	var readFunc func() BodyResponse
	headers, payload := SerializeHeaders(method, body), []byte{}

	switch method {
	case METHOD_GET, METHOD_POST, METHOD_PUT, METHOD_PATCH, METHOD_DELETE:
		payload = append(payload, body.Data...)
		withPayload = true
	}
	bodyRequest := append(headers, payload...)
	
	if wrted, err := conn.ref.Write(bodyRequest); wrted != len(bodyRequest) || err != nil {
		panic(errWarn(SENDREQ_ERR, conn.host))
	}

	if conn.file != "" && withPayload {
		readFunc = conn.syncReadResp
	} else {
		readFunc = conn.readResp
	}

	if conn.timeout > 0 { // if timeout is not set bypass deadline
		conn.ref.SetReadDeadline(time.Now().Add(conn.timeout))
	}
	response = readFunc()
	return conn.errormsg == nil, response
}

func (conn *linkc) SetFile(file string) {
	if _, err := os.Stat(file); err == nil {
		if err := os.Truncate(file, 0); err != nil {
			panic(errWarn(ERASEFILE_ERR, file))
		}
	}
	conn.file = file
}

func (conn *linkc) GetError() error {
	return conn.errormsg
}

func (conn *linkc) Quit() {
	conn.ref.Close()
	close(conn.ContentLength)
	close(conn.ContentDownloaded)
}
