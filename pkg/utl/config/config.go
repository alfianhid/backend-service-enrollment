package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server    *Server      `yaml:"server,omitempty"`
	App       *Application `yaml:"application,omitempty"`
	DB        *Database    `yaml:"database,omitempty"`
	Keycloaks *Keycloaks   `yaml:"keycloaks,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
}

// Database holds data necessery for database configuration
type Database struct {
	Dialect  string `yaml:"dialect,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Name     string `yaml:"name,omitempty"`
	Protocol string `yaml:"protocol,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Settings string `yaml:"settings,omitempty"`
}

type Keycloaks struct {
	Realm        string `yaml:"realm,omitempty"`
	ClientSecret string `yaml:"client_secret,omitempty"`
	ClientId     string `yaml:"cliend_id,omitempty"`
	Server       string `yaml:"server,omitempty"`
	UserAdmin    string `yaml:"user_admin,omitempty"`
	PassAdmin    string `yaml:"pass_admin,omitempty"`
	RealmAdmin   string `yaml:"realm_admin,omitempty"`
}

// LoadConfigFrom returns Configuration struct compile from input path
// reads the input file and builds a config struct
// that is serialized from all the data in the config rile
func LoadConfigFrom(cfgPath string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %v", err)
	}
	var cfg = new(Configuration)
	if err = yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config yaml into struct, %v", err)
	}
	return cfg, nil
}
