<h1 align="center">downbuff</h1>

<p align="center">An easy-use raw HTTP request library</p>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/ANDRVV/downbuff)](https://goreportcard.com/report/github.com/ANDRVV/downbuff) [![Go](https://github.com/ANDRVV/downbuff/actions/workflows/go.yml/badge.svg)](https://github.com/ANDRVV/downbuff/actions/workflows/go.yml)

</div>

## Request page example

```go
package main

import (
	"fmt"
	"time"
	"github.com/ANDRVV/downbuff"
)

func main() {
	headers := downbuff.Header{
		USER_AGENT:      "downbuff/1.0",
		ACCEPT:          "text/html",
		ACCEPT_LANGUAGE: "en-US,en;q=0.8",
		HOST:            "google.com",
	}

	baseRequest := downbuff.BodyRequest{Path: "/", HttpVersion: "1.1", Header: headers}

	conn := downbuff.Link("https://google.com", downbuff.SockConf{Buffer: 2048, Timeout: 3 * time.Second})
    	defer conn.Quit()

	ok, response1 := conn.Req(downbuff.METHOD_GET, baseRequest)
	if !ok {
		fmt.Println("error:", conn.GetError().Error())
	}
	fmt.Println(response1.StatusCode == downbuff.CODE_200)

	baseRequest.Path = "/maps"
	ok, response2 := conn.Req(downbuff.METHOD_GET, baseRequest)
	if !ok {
		fmt.Println("error:", conn.GetError().Error())
	}
	fmt.Println(response2.StatusCode == downbuff.CODE_200)

	fmt.Println(response1.HttpVersion, response1.StatusCode, response1.StatusText) // print HTTP version, status code, status text
	fmt.Println(response1.Summary())                                               // print serialized known headers + unknown headers
	fmt.Println(response1.Data)                                                    // print body of response
	fmt.Println(response1.Header.CONTENT_ENCODING)                                 // print stuff

	// ... same for response2
}
```

## Download content example

```go
package main

import (
	"fmt"
	"github.com/ANDRVV/downbuff"
	"time"
)

func main() {
	headers := downbuff.Header{
		RANGE: "bytes=0-", // Setting range, show content-length (good for performance in sync downloading)
		HOST:  "esahubble.org",
	}

	baseRequest := downbuff.BodyRequest{Path: "/media/archives/images/publicationtiff40k/heic1502a.tif", HttpVersion: "1.1", Header: headers}

	conn := downbuff.Link("https://esahubble.org", downbuff.SockConf{Buffer: 2048}) // remove timeout param for download big data
	defer conn.Quit()

	conn.SetFile("spaceimage.tif") // enable sync downloading with chan (ContentLength, ContentDownloaded)

	go func() {
		var contentLength int64
		var startTime time.Time
		for {
			select {
			case cl := <-conn.ContentLength:
				contentLength = cl
				startTime = time.Now()
			case cd := <-conn.ContentDownloaded:
				downloadSpeed := float64(cd) / time.Since(startTime).Seconds() / 1e6
				remainingTime := time.Duration((((contentLength - cd) / 1e6) / int64(max(downloadSpeed, 1))) * int64(time.Second)).String()

				fmt.Printf("\rDownload progress: %.2f GB / %.2f GB (%.2f MB/s): %s remaining time...", float64(cd)/1e9, float64(contentLength)/1e9, downloadSpeed, remainingTime)
				if contentLength == cd {
					return
				}
			}
		}
	}()

	ok, response := conn.Req(downbuff.METHOD_GET, baseRequest)
	fmt.Println(response.StatusCode)
	if ok && response.StatusCode == downbuff.CODE_200 {
		fmt.Println("Downloaded successfully")
	}
}
```

## Build Auth and Post

```go
package main

import (
	"github.com/ANDRVV/downbuff"
)

func main() {
	// How to use Auth builder

	credBasic := downbuff.Basic{Username: "user", Password: "passwd"}

	// Algorithm param not needed but easily assignable (default is MD5)
	credDigest := downbuff.Digest{
		Method: downbuff.METHOD_GET, 
		Path: "/", 
		Username: "user", 
		Password: "passwd",
		Algorithm: downbuff.MD5,
	}

	_ = downbuff.Header{
		// Auth with "Basic"
		AUTHORIZATION: downbuff.BuildAuth(downbuff.BASIC, downbuff.AuthInfo{BasicScheme: credBasic}),
	}

	_ = downbuff.Header{
		// Auth with "Digest"
		AUTHORIZATION: downbuff.BuildAuth(downbuff.DIGEST, downbuff.AuthInfo{DigestScheme: credDigest}),
	}

	// How to build data for POST request

	baseRequest := downbuff.BodyRequest{Path: "/", HttpVersion: "1.1", Header: downbuff.Header{HOST: "example.com"},
		Data: downbuff.BuildPOST(map[string]string{"key1": "value1", "key2": "value2"}),
	}

	conn := downbuff.Link("https://example.com", downbuff.SockConf{})
	defer conn.Quit()

	conn.Req(downbuff.METHOD_POST, baseRequest)
	
	// some stuff here
}
```

> [!NOTE]
> To add multiple or unknown headers use `UnkHeader` from `downbuff.BodyRequest` with `map[string]string` syntax.

<h2 align="center">Credits</h2>

<p align="center">Andrea Vaccaro</p>

<hr>

<p>Docs and update soon🚀</p>
