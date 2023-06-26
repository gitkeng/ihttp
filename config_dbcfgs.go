package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
)

type DBConfigs struct {
	IConfig
	Databases []*DBConfig `mapstructure:"databases" json:"databases"`
}

func (db *DBConfigs) Bind() error {
	for _, dbConf := range db.Databases {
		if err := dbConf.Bind(); err != nil {
			return err
		}
	}
	return nil
}

func (db *DBConfigs) Validate() error {
	checkDuplicate := make(map[string]bool)
	for _, dbConf := range db.Databases {
		if err := dbConf.Validate(); err != nil {
			return err
		}
		if _, ok := checkDuplicate[dbConf.ContextName]; ok {
			return ErrDuplicateDBContextName(dbConf.ContextName)
		} else {
			checkDuplicate[dbConf.ContextName] = true
		}
	}
	return nil
}

func (db *DBConfigs) String() string {
	return stringutil.Json(*db)
}

func (db *DBConfigs) ToMap() map[string]any {
	return convutil.Obj2Map(*db)
}
