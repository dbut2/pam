package server

type Config struct {
	Address string `yaml:"address"`
	Prod    bool
}
