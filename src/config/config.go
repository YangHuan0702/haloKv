package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ServerConfig struct {
	Port int32  `yaml:"Port"`
	Name string `yaml:"Name"`
}

type Raft struct {
	Servers          string `yaml:"Servers"`
	ElectTime        int32  `yaml:"ElectTime"`
	HeartbeatTimeOut int32  `yaml:"HeartbeatTimeOut"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
	Raft         Raft         `yaml:"Raft"`
	LogPath      string       `yaml:"LogPath"`
}

func ReadConfig() Config {
	file, err := os.ReadFile("config/application.yml")
	if err != nil {
		panic("read config/application.yml configuration file panic!")
	}
	config := Config{}
	if err = yaml.Unmarshal(file, &config); err != nil {
		panic(err.Error())
	}
	return config
}
