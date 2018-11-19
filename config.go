package configfacade

// Config is the interface to implement to make compatible library
type Config interface {
	LoadFile(path string, filename string, extension string) error
	LoadEnvVars(vars map[string]string) error
	Get(key string) interface{}
}

// Settings is a structure containing the main information about the configuration file and environment variables
type Settings struct {
	Path      string
	Name      string
	Extension string
	EnvVars   map[string]string
}

// Init initialises the library instance with settings
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
