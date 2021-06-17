package main

import (
	"log"

	"github.com/daheige/cobra-demo/cmd"
)

func main() {
	log.Println("cobra demo...")
	cmd.InitConfig()
	cmd.Execute()
}
