package main

import (
	"main.go/config"
	"main.go/router"
)

var (
	logger config.Logger
)

//@title WebNewsAPI
func main(){

	logger = *config.GetLogger("main")
  	//Initialzie Configs
	err := config.Init()
	if(err != nil){
		logger.Errorf("config initializarion error: %v", err)
		return
	}

	router.Initialize()
}
