package main

import (
	"errors"
)

type warnlog string

const (
	SOCKET_ERR              warnlog = "Unable to initialize socket"
	SENDREQ_ERR             warnlog = "Unable to send any request"
	RECVRESP_ERR            warnlog = "Unable to receive any response"
	INVALID_HTTPVERSION_ERR warnlog = "Invalid HTTP version"
	INVALID_METHOD_ERR      warnlog = "Invalid HTTP method"
	OVERFLOW_MEM_ERR        warnlog = "Data too large to store in memory: use file method"
	INVALID_URL_ERR         warnlog = "Invalid URL syntax"
	INVALID_SCHEME_ERR      warnlog = "Invalid scheme"
	ERASEFILE_ERR           warnlog = "Unable to erase file"
	INVALID_AUTHSCHEME_ERR  warnlog = "Invalid authentication scheme"
	PARAM_NEEDED_ERR        warnlog = "Needed parameter is missing"
	INVALID_HASH_ERR        warnlog = "Invalid hash algorithm"
)

func errWarn(log warnlog, extra string) (warnmsg error) {
	if extra == "" {
		return errors.New("<ERROR> " + string(log) + ".")
	}
	return errors.New("<ERROR> " + string(log) + ": " + extra)
}
