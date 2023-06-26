package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
)

type RedisConfigs struct {
	IConfig
	Caches []*RedisConfig `mapstructure:"redis" json:"redis"`
}

func (c *RedisConfigs) Bind() error {
	for _, cache := range c.Caches {
		if err := cache.Bind(); err != nil {
			return err
		}
	}
	return nil
}

func (c *RedisConfigs) Validate() error {
	checkDuplicate := make(map[string]bool)
	for _, cache := range c.Caches {
		if err := cache.Validate(); err != nil {
			return err
		}
		if _, ok := checkDuplicate[cache.ContextName]; ok {
			return ErrDuplicateRedisContextName(cache.ContextName)
		} else {
			checkDuplicate[cache.ContextName] = true
		}
	}
	return nil
}

func (c *RedisConfigs) String() string {
	return stringutil.Json(*c)
}

func (c *RedisConfigs) ToMap() map[string]any {
	return convutil.Obj2Map(*c)
}
