package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	configStore ConfigService

	// errors
	ErrUnableToLoad = errors.New("unable to load config")
)

type ConfigService interface {
	GetJWTSecret() string
	GetDBPath() string
	Error() error
	Load(location string) error
}

// If no configuration exists in memory: loads configuration from a file and stores it in memory at a fixed location
// always: returns a pointer to configuration in memory
func ConfigurationService(location string) ConfigService {
	if configStore == nil {
		cfg := new(config)
		err := cfg.Load(location)

		configStore = cfg

		cfg.error = err

		return cfg
	}

	return configStore
}

type config struct {
	JWTSecret string `yaml:"jwt_secret"`
	DBPath    string `yaml:"db_path"`

	error error

	mutex sync.RWMutex
}

func (c *config) GetDBPath() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.DBPath
}

// Concurrency safe fetching of the JWT secret
func (c *config) GetJWTSecret() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.JWTSecret
}

// Load a config from a location on disk
func (c *config) Load(location string) error {
	_, err := os.Stat(location)
	// config file location doesn't exist
	if os.IsNotExist(err) {
		// create a new file
	}
	if !os.IsNotExist(err) && err != nil {
		return fmt.Errorf("%s: %w", ErrUnableToLoad, err)
	}

	body, err := ioutil.ReadFile(location)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrUnableToLoad, err)
	}

	return yaml.NewDecoder(bytes.NewReader(body)).Decode(c)
}

// Used for detecting errors after loading
func (c *config) Error() error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.error
}
