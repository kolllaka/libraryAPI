package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/KoLLlaka/libraryAPI/internal/apiserver"
)

var configPath = "configs/apiserver.toml"

func main() {
	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
