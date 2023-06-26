package dbutil_test

import (
	"fmt"
	"testing"

	"github.com/gitkeng/ihttp/util/dbutil"
)

func TestSQLLoader(t *testing.T) {
	fileName := "thai-division.sql"
	sqlCmd, err := dbutil.SQLLoader(fileName)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	fmt.Printf("%s", sqlCmd)

}
