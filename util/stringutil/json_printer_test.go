package stringutil_test

import (
	"github.com/gitkeng/ihttp/util/stringutil"
	"testing"
)

func TestPrintJson(t *testing.T) {
	testData := struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Age      int    `json:"age"`
	}{
		Name:     "gitkeng",
		LastName: "Limpitikeat",
		Age:      20,
	}
	t.Logf("%s", stringutil.Json(testData))
	t.Logf("%s", stringutil.JsonPretty(testData))
}
