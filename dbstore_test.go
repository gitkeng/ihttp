package ihttp_test

import (
	"github.com/gitkeng/ihttp"
	"testing"
)

func TestNewDBStore(t *testing.T) {
	configFileLocation := "./examples/conf/dbconfig.yaml"
	conf, err := ihttp.NewConfig(configFileLocation)
	if err != nil {
		t.Log(err)
		return
	}
	mysqlDBCfg, found := conf.GetDBConfig("mysqldb")
	if !found {
		t.Errorf("Database context name mysqldb Not found")
		return
	}

	mysqldbStore, err := ihttp.NewDBStore(mysqlDBCfg)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("mysqldbStore config: %s", mysqldbStore.Config())
	err = mysqldbStore.Close()
	if err != nil {
		t.Log(err)
	}

	pqDBCfg, found := conf.GetDBConfig("pgdb")
	if !found {
		t.Errorf("Database context name pgdb Not found")
		return
	}

	pqdbStore, err := ihttp.NewDBStore(pqDBCfg)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("pqdbStore config: %s", pqdbStore.Config())
	err = pqdbStore.Close()
	if err != nil {
		t.Log(err)
	}
}
