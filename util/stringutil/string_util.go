package stringutil

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

// IsEmptyString return true if empty string
func IsEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

// IsNotEmptyString return true if not empty string
func IsNotEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) > 0
}

func Compress(data string, isEncode bool) (string, error) {
	byteData := []byte(data)
	buff := &bytes.Buffer{}
	gz := gzip.NewWriter(buff)
	_, err := gz.Write(byteData)
	if err != nil {
		return "", err
	}
	err = gz.Flush()
	if err != nil {
		return "", err
	}
	err = gz.Close()
	if err != nil {
		return "", err
	}
	if isEncode {
		compressString := base64.StdEncoding.EncodeToString(buff.Bytes())
		return compressString, nil
	} else {
		return buff.String(), nil
	}
}

func DeCompress(compressData string, isEncode bool) (string, error) {
	var reader *bytes.Reader
	if isEncode {
		compressByte, err := base64.StdEncoding.DecodeString(compressData)
		if err != nil {
			return "", err
		}
		reader = bytes.NewReader(compressByte)
	} else {
		reader = bytes.NewReader([]byte(compressData))
	}
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(gzReader)
	if err != nil {
		return "", err
	}
	return string(data), nil

}
