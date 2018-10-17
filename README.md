# Config Facade
[![Build Status](https://travis-ci.org/limoli/configfacade.svg?branch=master)](https://travis-ci.org/limoli/configfacade)

Config Facade is an useful facade for multiple configurations using different libraries (e.g. viper).

# Problem
Many times we begin projects with different libraries in order to provide configurations for our apps and every time we have to implement new logics for the same features. This is a **big problem of abstraction**.
When you decide to use a specific library, you import it from github and you use it for whole project. But what happens if you need to change library for some reasons? You have to refactor everything and implement again the logics if you are enough unlucky.

# Solution
It would be nice if a common wrapper could implement the main features without thinking too much about which libraries we have decided to use. In this way, we will be able to use different libraries in the future and extend our projects compatibility.

What you will able to do:
- set a complete configuration in few lines of code
- choose any compatible library
 
# Compatible libraries
- [Viper](https://github.com/spf13/viper)

# Installation

```sh
dep ensure --add github.com/limoli/configfacade
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
```go
package main

import (
    "github.com/limoli/configfacade"
    "github.com/limoli/configfacade/viper"
    "log"
    "os"
)

var Config config.Config

func main() {
    // Initialisation 
    Config, err := config.Init(new(viper.Config), config.Settings{
        Path:      "./extra/config/",
        Name:      os.Getenv("APP_ENV"),
        Extension: "yaml",
        EnvVars: []config.EnvVar{
            {"app.port", "APP_PORT"},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // Use of configuration
    port := Config.Get("app.port").(string)
    log.Println("Port", port)

}
``` 

# Contribute
Would you like to see a library in the compatible list? 
Just ask or implement yourself the `Config` interface and send a pull request.

```go
type Config interface {
	LoadFile(path string, filename string, extension string) error
	LoadEnvVars(vars []EnvVar) error
	Get(key string) interface{}
}
```



