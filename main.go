package main

import (
	"fmt"
	"log"

	"github.com/takama/daemon"
)

func main() {
	service, err := daemon.New("name_1", "desc_1", daemon.SystemDaemon)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	status, err := service.Install()
	if err != nil {
		log.Fatal(status, "\nError: ", err)
	}

	fmt.Println(status)
}
