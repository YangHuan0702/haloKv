package config

import (
	"gopkg.in/yaml.v2"
	"haloKv/src/common"
	"os"
)

type ServerConfig struct {
	Port int32  `yaml:"Port"`
	Name string `yaml:"Name"`
}

type RaftConfig struct {
	Servers          string `yaml:"Servers"`
	ElectTime        int    `yaml:"ElectTime"`
	HeartbeatTimeOut int    `yaml:"HeartbeatTimeOut"`
	Me               int    `yaml:"Me"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
	Raft         RaftConfig   `yaml:"Raft"`
	LogPath      string       `yaml:"LogPath"`
}

func ReadConfig() Config {
	file, err := os.ReadFile(common.ConfigPath)
	if err != nil {
		panic("read config/application.yml configuration file panic!")
	}
	config := Config{}
	if err = yaml.Unmarshal(file, &config); err != nil {
		panic(err.Error())
	}
	return config
}
