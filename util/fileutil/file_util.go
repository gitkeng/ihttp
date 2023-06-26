package fileutil

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func LineCounter(fileLocation string) (int, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}

func IsFileExist(fileLocation string) (bool, error) {
	fi, err := os.Stat(fileLocation)
	if err != nil {
		return false, err
	}
	mode := fi.Mode()
	return mode.IsRegular(), nil
}

func IsDirExist(dirLocation string) (bool, error) {
	fi, err := os.Stat(dirLocation)
	if err != nil {
		return false, err
	}
	mode := fi.Mode()
	return mode.IsDir(), nil
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsTextFile(fileLocation string) (bool, error) {
	mime, err := mimetype.DetectFile(fileLocation)
	if err != nil {
		return false, err
	}
	return strings.Contains(mime.String(), "text/plain"), nil
}

func TextFileLoader(fileLocation string) (string, error) {
	if isTextFile, _ := IsTextFile(fileLocation); isTextFile {
		f, err := os.Open(fileLocation)
		if err != nil {
			return "", err
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
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
	return "", fmt.Errorf("%s is not text file", fileLocation)

}

func IsImageFile(fileLocation string) (bool, error) {
	mime, err := mimetype.DetectFile(fileLocation)
	if err != nil {
		return false, err
	}
	return strings.Contains(mime.String(), "image/"), nil
}

func ReadLine(fileLocation string, lineNum int) (line string, lastLine int, err error) {
	if lineNum < 1 {
		return "", 0, fmt.Errorf("read line invalid request: line %d", lineNum)
	}

	file, err := os.Open(fileLocation)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

func GetFileSize(fileLocation string) (int64, error) {
	fi, err := os.Stat(fileLocation)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

func AppendCSV(fileLocation string, csvData []string, addNewLine bool) error {
	f, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if addNewLine {
		f.WriteString("\n")
	}

	w := csv.NewWriter(f)
	err = w.Write(csvData)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func Read(fileLocation string) ([]byte, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func ReadLines(fileLocation string) ([]string, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteLines writes the lines to the given file.
func WriteLines(lines []string, fileLocation string) error {
	file, err := os.Create(fileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// Write writes the binary to the given file.
func Write(b []byte, fileLocation string) (int, error) {
	file, err := os.Create(fileLocation)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Write(b)
}

func MoveFile(currentLocation string, newLocation string) error {
	isExist, _ := IsFileExist(newLocation)
	if isExist {
		return fmt.Errorf("%s is exist", newLocation)
	}

	err := os.Rename(currentLocation, newLocation)
	if err != nil {
		return err
	}

	return nil

}

func CopyFile(source string, destination string) error {
	isExist, _ := IsFileExist(destination)
	if isExist {
		return fmt.Errorf("%s is exist", destination)
	}

	from, err := os.Open(source)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}
	return nil
}

func IsNewLineAtEOF(fileLocation string) (bool, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return false, err
	}
	defer file.Close()
	buf := make([]byte, 10)
	lastBuf := make([]byte, 10)
	lastByteRead := 0
	for {
		byteread, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}
		lastBuf = buf
		lastByteRead = byteread
	}
	lastReadString := string(lastBuf[:lastByteRead])
	lastString := string(lastReadString[len(lastReadString)-1])

	return lastString == "\n", nil
}

func DeleteFile(fileLocation string) error {
	ok, err := IsFileExist(fileLocation)
	if err == nil {
		if !ok {
			return fmt.Errorf("%s not file", fileLocation)
		}
		err = os.Remove(fileLocation)
		if err != nil {
			return err
		}
		return nil
	}
	return err

}

func DeleteFolder(dirLocation string) error {
	ok, err := IsDirExist(dirLocation)
	if err == nil {
		if !ok {
			return fmt.Errorf("%s not directory", dirLocation)
		}
		err = os.Remove(dirLocation)
		if err != nil {
			return err
		}
		return nil

	}
	return err
}

func GetDir(fileLocation string) (string, error) {
	p, err := filepath.Abs(fileLocation)
	if err != nil {
		return "", err
	}
	return filepath.Dir(p), nil
}

func GetFileNameWithExt(fileLocation string) (string, error) {
	p, err := filepath.Abs(fileLocation)
	if err != nil {
		return "", err
	}
	return filepath.Base(p), nil
}

func GetFileNameOnly(fileLocation string) (string, error) {
	p, err := filepath.Abs(fileLocation)
	if err != nil {
		return "", err
	}
	fileNameWithExt := filepath.Base(p)
	extension := filepath.Ext(p)
	return fileNameWithExt[0 : len(fileNameWithExt)-len(extension)], nil
}

func GetFileExtension(fileLocation string) (string, error) {
	p, err := filepath.Abs(fileLocation)
	if err != nil {
		return "", err
	}
	extension := filepath.Ext(p)
	return extension, nil
}

func ListFileInDirectory(rootDir string, isFindSubDir bool) ([]string, error) {
	lastDirString := rootDir[len(rootDir)-1:]
	if lastDirString != string(os.PathSeparator) {
		rootDir = rootDir + string(os.PathSeparator)
	}

	var files []string
	err := filepath.Walk(rootDir, visit(&files, rootDir, isFindSubDir))
	if err != nil {
		return []string{}, err
	}
	return files, nil
}

func visit(files *[]string, rootDir string, isFindSubDir bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if filepath.Dir(rootDir) == filepath.Dir(path) && !isFindSubDir {
				if string(filepath.Base(path)[0]) != "." {
					*files = append(*files, path)
				}
			}

			if isFindSubDir {
				if string(filepath.Base(path)[0]) != "." {
					*files = append(*files, path)
				}
			}

		}
		return nil
	}
}
