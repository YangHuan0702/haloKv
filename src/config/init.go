package config

import "fmt"

func init() {
	config := ReadConfig()
	fmt.Println("init config:", config)
}
