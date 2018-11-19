package configfacade

type Config interface {
	LoadFile(path string, filename string, extension string) error
	LoadEnvVars(vars map[string]string) error
	Get(key string) interface{}
}

type Settings struct {
	Path      string
	Name      string
	Extension string
	EnvVars   map[string]string
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
