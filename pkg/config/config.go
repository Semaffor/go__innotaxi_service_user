package config

import "time"

type ServerConfig struct {
	Host           string        `mapstructure:"host" default:"localhost" json:"host,omitempty"` // TODO: default not working
	Port           string        `mapstructure:"port" default:"8000" json:"port,omitempty"`
	MaxHeaderBytes int           `mapstructure:"maxHeaderBytes" default:"20"` // number of bit shifts to the left
	ReadTimeout    time.Duration `mapstructure:"readTimeout" default:"10"`    // time in seconds
	WriteTimeout   time.Duration `mapstructure:"writeTimeout" default:"10"`   // time in seconds
}

type ConfigDb struct {
	Host     string `mapstructure:"host" default:"localhost"`
	Port     string `mapstructure:"port" default:"8001"`
	Username string `mapstructure:"username"`
	DbName   string `mapstructure:"dbName"`
	SslMode  string `mapstructure:"sslMode"`
	Password string `mapstructure:"password"`
}
