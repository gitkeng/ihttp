package dbutil

import (
	"bufio"
	"bytes"
	"os"
)

func SQLLoader(fileLocation string) (string, error) {
	fSQL, err := os.Open(fileLocation)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(fSQL)
	defer fSQL.Close()
	done := make(chan struct{})
	buffer := &bytes.Buffer{}
	go func() {
		defer close(done)
		for scanner.Scan() {
			buffer.WriteString(scanner.Text() + "\n")
		}
	}()
	<-done
	return buffer.String(), nil

}
