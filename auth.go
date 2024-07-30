package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type AuthScheme string
type DigestAlgorithm string

type Basic struct {
	Username, Password string
}

type Digest struct {
	Method reqmethod // needed
	Path   string    // default is "/"

	Username   string
	Password   string
	Realm      string
	Algorithm  DigestAlgorithm
	LastNonce  string
	Nonce      string
	NonceValue int
	Qop        string
	Opaque     string
}

type AuthInfo struct {
	BasicScheme  Basic
	DigestScheme Digest
}

const (
	// Authentication Scheme
	BASIC  AuthScheme = "Basic"
	DIGEST AuthScheme = "Digest"

	// Supported algorithms
	MD5     DigestAlgorithm = "MD5"
	MD5SESS DigestAlgorithm = "MD5-SESS"
	SHA     DigestAlgorithm = "SHA"
	SHA256  DigestAlgorithm = "SHA-256"
	SHA512  DigestAlgorithm = "SHA-512"
)

// Value for Proxy-Authorization header
func BuildProxyAuth(info Basic) string {
	return buildBasicAuth(info)
}

// Value for Authorization header
func BuildAuth(scheme AuthScheme, authInfo AuthInfo) (headerContent string) {
	switch scheme {
	case BASIC:
		return buildBasicAuth(authInfo.BasicScheme)
	case DIGEST:
		return buildDigestAuth(authInfo.DigestScheme)
	default:
		panic(errWarn(INVALID_AUTHSCHEME_ERR, string(scheme)))
	}
}

func buildBasicAuth(info Basic) string {
	login := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", info.Username, info.Password)))
	return fmt.Sprintf("Basic %s", login)
}

func buildDigestAuth(info Digest) string {
	if info.Method == "" {
		panic(errWarn(PARAM_NEEDED_ERR, "info.Method is needed in Digest builder"))
	}
	if info.Path == "" {
		info.Path = "/" // default
	}

	algHash := func(s string) string {
		var hash []byte
		switch info.Algorithm {
		case MD5, MD5SESS, "":
			sum := md5.Sum([]byte(s))
			hash = sum[:]
		case SHA:
			sum := sha1.Sum([]byte(s))
			hash = sum[:]
		case SHA256:
			sum := sha256.Sum256([]byte(s))
			hash = sum[:]
		case SHA512:
			sum := sha512.Sum512([]byte(s))
			hash = sum[:]
		default:
			panic(errWarn(INVALID_HASH_ERR, string(info.Algorithm)))
		}
		return hex.EncodeToString(hash[:])
	}

	HashA1 := algHash(fmt.Sprintf("%s:%s:%s", info.Username, info.Realm, info.Password))
	HashA2 := algHash(fmt.Sprintf("%s:%s", info.Method, info.Path))

	if info.Nonce == info.LastNonce {
		info.NonceValue++
	} else {
		info.NonceValue = 1
	}

	s := strconv.Itoa(info.NonceValue) + info.Nonce + time.Now().Format(time.ANSIC)
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	s += string(randomBytes)

	ncHexValue := fmt.Sprintf("%08x", max(info.NonceValue, 1))
	cnonce := algHash(s)[:16]

	if info.Algorithm == MD5SESS {
		HashA1 = algHash(fmt.Sprintf("%s:%s:%s", HashA1, info.Nonce, cnonce))
	}

	var respContent string
	if info.Qop == "" {
		respContent = algHash(fmt.Sprintf("%s:%s:%s", HashA1, info.Nonce, HashA2))
	} else if info.Qop == "auth" || slices.Contains(strings.Split(info.Qop, ","), "auth") {
		respContent = algHash(fmt.Sprintf("%s:%s:%s:%s:auth:%s", HashA1, info.Nonce, ncHexValue, cnonce, HashA2))
	} else {
		return ""
	}

	base := fmt.Sprintf("username=\"%s\", realm=\"%s\", nonce=\"%s\", uri=\"%s\", response=\"%s\"", info.Username, info.Realm, info.Nonce, info.Path, respContent)

	if info.Algorithm != "" {
		base += ", algorithm=\"%s\""
	}
	if info.Opaque != "" {
		base += fmt.Sprintf(", opaque=\"%s\"", info.Opaque)
	}
	if info.Qop != "" {
		base += fmt.Sprintf(", qop=\"auth\", nc=%s, cnonce=\"%s\"", ncHexValue, cnonce)
	}

	return fmt.Sprintf("Digest %s", base)
}
