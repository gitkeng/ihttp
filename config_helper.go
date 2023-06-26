package ihttp

import (
	"github.com/gitkeng/ihttp/util/fileutil"
	"github.com/spf13/viper"
	"strings"
)

func ReadConfigFile[T IConfig](fileLocation string, config T) error {
	vp := viper.New()
	vp.AutomaticEnv()
	path, err := fileutil.GetDir(fileLocation)
	if err != nil {
		return err
	}
	fileName, err := fileutil.GetFileNameOnly(fileLocation)
	if err != nil {
		return err
	}
	vp.SetConfigName(fileName)
	vp.AddConfigPath(path)

	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	err = vp.Unmarshal(config)
	if err != nil {
		return err
	}

	if err := config.Bind(); err != nil {
		return err
	}

	if err := config.Validate(); err != nil {
		return err
	}

	return nil
}

func Getenv(envName string, prefix ...string) string {
	if len(prefix) > 0 {
		vp := viper.New()
		vp.SetEnvPrefix(strings.Join(prefix, "_"))
		vp.AutomaticEnv()
		return vp.GetString(envName)
	} else {
		vp := viper.New()
		vp.AutomaticEnv()
		return vp.GetString(envName)
	}
}
