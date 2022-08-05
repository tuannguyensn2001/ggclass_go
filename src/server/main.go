package main

import (
	"ggclass_go/src/cmd"
	"ggclass_go/src/config"
	"ggclass_go/src/packages/logger"
	"log"
)

func main() {

	err := config.Load()

	logger.InitLog()

	defer config.Cfg.GetRabbitMQ().Close()

	if err != nil {
		log.Fatalln("err load config", err)
	}

	rootCmd := cmd.GetRoot()

	err = rootCmd.Execute()

	if err != nil {
		log.Fatalln("err command", err)
	}

}
