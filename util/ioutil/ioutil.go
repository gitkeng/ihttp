package ioutil

import (
	"bytes"
	"io"
)

func Byte2Reader(b []byte) io.Reader {
	return bytes.NewReader(b)
}

func Byte2ReaderSeeker(b []byte) io.ReadSeeker {
	return bytes.NewReader(b)
}
