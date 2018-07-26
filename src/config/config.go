package config

import (
	"log"
	"../github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
// c is a pointer, hence reading into it will update the value for the caller also
// Read is a function attached to this Config struct -> caller only needs to call Config.Read() to init the struct
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}