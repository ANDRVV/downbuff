package main

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

func checkValid(kv map[string]string) (err error) {
	for key := range kv {
		switch key {
		case "REQ_METHOD":
			if containsLeast(kv[key], []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE", "CONNECT"}) {
				continue
			}
			return errWarn(INVALID_METHOD_ERR, kv[key])
		case "REQ_HTTPVERSION":
			if containsLeast(kv[key], []string{"1.0", "1.1", "2", "3"}) {
				continue
			}
			return errWarn(INVALID_HTTPVERSION_ERR, kv[key])
		}
	}
	return nil
}

func containsLeast(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func toValidAddr(rawUrl string) (validUrl string, scheme string, err error) {
	if b, _, f := strings.Cut(rawUrl, "://"); f {
		scheme = strings.ToLower(b)
		if scheme == "" {
			scheme = "http"
		} else if !slices.Contains([]string{"http", "https"}, scheme) {
			panic(errWarn(INVALID_SCHEME_ERR, scheme))
		}
	} else {
		rawUrl = fmt.Sprintf("http://%s", rawUrl)
	}
	if pUrl, err := url.Parse(rawUrl); err == nil {
		host, port := pUrl.Hostname(), pUrl.Port()
		if port == "" {
			if scheme == "https" {
				port = "443"
			} else if scheme == "http" {
				port = "80"
			}
		}
		return fmt.Sprintf("%s:%s", host, port), scheme, nil
	}
	return "", "", errWarn(INVALID_URL_ERR, rawUrl)
}
