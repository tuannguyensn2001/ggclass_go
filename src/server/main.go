package main

import (
	"fmt"
	"ggclass_go/src/cmd"
	"ggclass_go/src/config"
	"log"
)

func main() {
	err := config.Load()

	fmt.Print("hello")

	if err != nil {
		log.Fatalln("err load config", err)
	}

	rootCmd := cmd.GetRoot()

	err = rootCmd.Execute()

	if err != nil {
		log.Fatalln("err command", err)
	}

}
