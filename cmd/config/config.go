package config

import (
	"os"

	"github.com/dbut2/pam/internal/app"
	"github.com/dbut2/pam/internal/server"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App    app.Config    `yaml:"app"`
	Server server.Config `yaml:"server"`
}

func Find() Config {
	c := findConfig()

	if creds, found := findCreds(); found {
		c.App.Auth.Creds = creds
	}

	return c
}

func findConfig() Config {
	c := Config{}

	_, err := os.Stat("config.yaml")
	if !os.IsNotExist(err) {
		if err != nil {
			panic(err.Error())
		}

		bytes, err := os.ReadFile("config.yaml")
		if err != nil {
			panic(err.Error())
		}

		err = yaml.Unmarshal(bytes, &c)
		if err != nil {
			panic(err.Error())
		}
	}

	_, err = os.Stat("/secrets/config.yaml")
	if !os.IsNotExist(err) {
		if err != nil {
			panic(err.Error())
		}

		bytes, err := os.ReadFile("/secrets/config.yaml")
		if err != nil {
			panic(err.Error())
		}

		err = yaml.Unmarshal(bytes, &c)
		if err != nil {
			panic(err.Error())
		}
	}

	return c
}

func findCreds() ([]byte, bool) {
	_, err := os.Stat("creds.json")
	if !os.IsNotExist(err) {
		if err != nil {
			panic(err.Error())
		}

		bytes, err := os.ReadFile("creds.json")
		if err != nil {
			panic(err.Error())
		}

		return bytes, true
	}

	_, err = os.Stat("/secrets/creds.json")
	if !os.IsNotExist(err) {
		if err != nil {
			panic(err.Error())
		}

		bytes, err := os.ReadFile("/secrets/creds.json")
		if err != nil {
			panic(err.Error())
		}

		return bytes, true
	}

	return nil, false
}
