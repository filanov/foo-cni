package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	logFile, err := os.OpenFile("/var/log/foo-cni.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("failed to open log file")
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// Parse command (ADD/DEL)
	cmd := os.Getenv("CNI_COMMAND")
	args := os.Getenv("CNI_ARGS")

	log.Println("CNI ARGS", args)

	switch cmd {
	case "ADD":
		log.Println("Setting up network...")
		// Add your network setup logic here
		// fmt.Fprintf(os.Stdout, `{"cniVersion": "%s", "interfaces": [{"name": "eth0"}]}`, config.CniVersion)
		fmt.Println("{}")
		os.Exit(0)
	case "DEL":
		log.Println("Cleaning up network...")
		// Add your network cleanup logic here
		fmt.Println("{}")
		os.Exit(0)
	default:
		log.Printf("Unsupported CNI_COMMAND: %s\n", cmd)
		os.Exit(1)
	}
}
