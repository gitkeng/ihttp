package fileutil_test

import (
	"github.com/gitkeng/ihttp/util/fileutil"
	"testing"
)

func TestGetDir(t *testing.T) {
	fileLocation := "./test/poc_file.txt"
	dir, err := fileutil.GetDir(fileLocation)
	if err != nil {
		t.Errorf("GetDir error: %s", err.Error())
	}
	t.Logf("%s", dir)
}

func TestGetFileNameOnly(t *testing.T) {
	fileLocation := "./test/poc_file.txt"
	fileName, err := fileutil.GetFileNameOnly(fileLocation)
	if err != nil {
		t.Errorf("GetFileNameOnly error: %s", err.Error())
	}
	t.Logf("%s", fileName)
}

func TestGetFileNameWithExt(t *testing.T) {
	fileLocation := "./test/poc_file.txt"
	fileName, err := fileutil.GetFileNameWithExt(fileLocation)
	if err != nil {
		t.Errorf("GetFileNameWithExt error: %s", err.Error())
	}
	t.Logf("%s", fileName)
}

func TestGetFileExtension(t *testing.T) {
	fileLocation := "./test/poc_file.txt"
	fileExt, err := fileutil.GetFileExtension(fileLocation)
	if err != nil {
		t.Errorf("GetFileExtension error: %s", err.Error())
	}
	t.Logf("%s", fileExt)
}
