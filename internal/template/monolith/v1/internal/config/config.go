package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"{{.Path}}/internal/errors"
)

const Filename = "{{.BinaryName}}.json"

// Settings contains the minimal bootstrapping configuration
type Settings struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Default returns the default configuration, which only makes sense for a developer machine, if anything at all.
func Default() Settings {
	return Settings{
		Host: "localhost",
		Port: 8080,
	}
}

// LoadFile decodes a json configuration from a file
func LoadFile(filename string) (s Settings, err error) {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return Default(), fmt.Errorf("config file '%s'cannot be opened: %w", filename, err)
	}

	defer errors.Check(file.Close, &err)

	r, err := Load(file)
	if err != nil {
		return Default(), fmt.Errorf("config file '%s' cannot be parsed: %w", filename, err)
	}

	return r, nil
}

// Load tries to decode the config from yaml format
func Load(reader io.Reader) (Settings, error) {
	r := Default()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(b, &r)

	return r, err
}
