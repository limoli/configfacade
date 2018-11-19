# Config Facade
[![GoDoc](https://godoc.org/limoli/dbshift?status.svg)](https://godoc.org/github.com/limoli/dbshift)
[![Go Report Card](https://goreportcard.com/badge/github.com/limoli/configfacade)](https://goreportcard.com/report/github.com/limoli/configfacade)
[![Maintainability](https://api.codeclimate.com/v1/badges/988c97ce3d1495e953a2/maintainability)](https://codeclimate.com/github/limoli/configfacade/maintainability)

Config Facade is a simple facade for multiple configurations using different libraries.

# Problem
Many times we begin projects with different libraries in order to provide configurations for our apps and every time we have to implement new logics for the same features. This is a **big problem of abstraction**.
When you decide to use a specific library, you import it from github and you use it for whole project. But what happens if you need to change library for some reason? You have to refactor everything and implement again the logics if you are enough unlucky.

# Solution
It would be nice if a common wrapper could implement the main features without thinking too much about which libraries we have decided to use. In this way, we will be able to use different libraries in the future and extend our projects compatibility.

What you will able to do:
- set a complete configuration in few lines of code
- choose any compatible library
 
# Compatible libraries
- [Viper](https://github.com/limoli/viper)

# Installation

```sh
dep ensure -add github.com/limoli/configfacade
``` 

# Example

**extra/config/development.yaml**
It defines a configuration file with `yaml` format for the development environment.
```yaml
app:
  port: "8080"
```

**main.go**
It initialises a configuration instance using Viper library. As for settings, it defines the configuration file to read and the possible environment variables which can override the default configuration value.

```sh
dep ensure -add github.com/limoli/viper@master
``` 

```go
import (
	"github.com/limoli/configfacade"
	"github.com/limoli/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var config configfacade.Config

func init() {
	var err error

	// filepath.Abs uses working directory and it can change during tests
	_, currentFilePath, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(currentFilePath), "config")

	config, err = configfacade.Init(new(viper.Facade), configfacade.Settings{
		Path:      path,
		Name:      os.Getenv("APP_ENV"),
		Extension: "yaml",
		EnvVars: map[string]string{
			"app.db.user":     "MYSQL_USER",
			"app.db.password": "MYSQL_PASSWORD",
			"app.db.host":     "MYSQL_HOST",
			"app.db.port":     "MYSQL_PORT",
			"app.db.name":     "MYSQL_DATABASE",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
``` 

# Contribute
Would you like to see a library in the compatible list? 
Just implement the `Config` interface and send a pull request.

```go
type Config interface {
	LoadFile(path string, filename string, extension string) error
	LoadEnvVars(vars map[string]string) error
	Get(key string) interface{}
}
```
The best thing you could do is to **implement this interface directly on the library**.


