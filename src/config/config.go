package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ServerConfig struct {
	Port int32
	Name string
}

type Raft struct {
	Servers          []string
	ElectTime        int32
	HeartbeatTimeOut int32
}

type Config struct {
	ServerConfig ServerConfig `json:"ServerConfig"`
	Raft         Raft         `json:"Raft"`
	LogPath      string       `json:"LogPath"`
}

func ReadConfig() Config {
	file, err := os.ReadFile("config/application.yml")
	if err != nil {
		panic("read config/application.yml configuration file panic!")
	}
	config := Config{}
	if err = yaml.Unmarshal(file, &config); err != nil {
		panic("yaml Unmarshal panic~")
	}
	return config
}
