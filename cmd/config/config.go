package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

//Config contains a link to a database connection, port settings
type Config struct {
	PORT        string `envconfig:"PORT" default:":7755"`
	User        string `envconfig:"USER" default:"postbox"`
	Pass        string `envconfig:"PASS" default:"master"`
	DBURL       string `envconfig:"DB_URL"`
	TemplateDir string `envconfig:"TEMP_DIR" default:"emails"`
}

//Init initializes config in App
func (c *Config) Init() {
	err := envconfig.Process("", c)
	if err != nil {
		log.Printf("Error %v when parse config\n", err)
		panic(err)
	}
	log.Printf("Config: %v\n", c)
}
