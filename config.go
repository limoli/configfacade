package configfacade

type Config interface {
	LoadFile(path string, filename string, extension string) error
	LoadEnvVars(vars []EnvVar) error
	Get(key string) interface{}
}

type EnvVar struct {
	Key string
	Env string
}

type Settings struct {
	Path      string
	Name      string
	Extension string
	EnvVars   []EnvVar
}

func Init(c Config, s Settings) (Config, error) {
	var err error

	err = c.LoadFile(s.Path, s.Name, s.Extension)
	if err != nil {
		return nil, err
	}

	err = c.LoadEnvVars(s.EnvVars)
	if err != nil {
		return nil, err
	}

	return c, nil
}
