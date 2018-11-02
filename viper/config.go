package viper

import (
	"github.com/limoli/configfacade"
	"github.com/spf13/viper"
)

type Config struct {
	instance *viper.Viper
}

func (c *Config) LoadFile(path, filename, extension string) error {
	c.instance = viper.New()
	c.instance.AddConfigPath(path)
	c.instance.SetConfigName(filename)
	c.instance.SetConfigType(extension)
	return c.instance.ReadInConfig()
}

func (c *Config) LoadEnvVars(vars []configfacade.EnvVar) error {
	var err error
	for _, v := range vars {
		err = c.instance.BindEnv(v.Key, v.Env)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) Get(key string) interface{} {
	return c.instance.Get(key)
}
