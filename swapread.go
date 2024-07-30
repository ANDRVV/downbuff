package downbuff

import (
	"bytes"
	"math"
	"os"
)

func (conn *linkc) syncReadResp() BodyResponse {
	file, _ := os.OpenFile(conn.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	response := BodyResponse{}
	chk := make([]byte, conn.buffer)
	var downloadedSize int64
	var data []byte

	for {
		if n, err := conn.ref.Read(chk); err == nil && n > 0 {
			data = append(data, chk[:n]...)
			downloadedSize += int64(n)
			if bytes.Contains(data, []byte{0x0d, 0x0a, 0x0d, 0x0a}) {
				response = packResponse(data)
				break
			}
		} else {
			break
		}
	}

	file.Write(response.Data)
	response.Data = []byte{}

	conn.ContentLength <- response.Header.CONTENT_LENGTH

	var toRead int64
	if response.Header.CONTENT_LENGTH > 1 {
		toRead = response.Header.CONTENT_LENGTH - downloadedSize
	} else {
		toRead = math.MaxInt64 - downloadedSize
	}

	for downloadedSize <= toRead {
		if n, err := conn.ref.Read(chk); err == nil && n > 0 {
			file.Write(chk[:n])
			downloadedSize += int64(n)
			conn.ContentDownloaded <- downloadedSize
		} else {
			break
		}
	}
	conn.ContentDownloaded <- response.Header.CONTENT_LENGTH

	return response
}

func (conn *linkc) readResp() BodyResponse {
	response := BodyResponse{}
	chk := make([]byte, conn.buffer)

	for {
		if len(response.Data)+conn.buffer >= math.MaxInt32 {
			conn.errormsg = errWarn(OVERFLOW_MEM_ERR, "")
			break
		} else {
			if n, err := conn.ref.Read(chk); err == nil && n > 0 {
				response.Data = append(response.Data, chk[:n]...)
			} else {
				break
			}
		}
	}
	return packResponse(response.Data)
}
