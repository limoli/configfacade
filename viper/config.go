package viper

import (
	"github.com/limoli/configfacade"
	"github.com/spf13/viper"
)

type Config struct{}

func (c *Config) Get(key string) interface{} {
	return viper.Get(key)
}

func (c *Config) LoadFile(path, filename, extension string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(extension)
	return viper.ReadInConfig()
}

func (c *Config) LoadEnvVars(vars []configfacade.EnvVar) error {
	var err error
	for _, v := range vars {
		err = viper.BindEnv(v.Key, v.Env)
		if err != nil {
			return err
		}
	}
	return nil
}
