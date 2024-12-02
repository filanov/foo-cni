package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse command (ADD/DEL)
	cmd := os.Getenv("CNI_COMMAND")
	args := os.Getenv("CNI_ARGS")

	fmt.Println("CNI ARGS", args)

	switch cmd {
	case "ADD":
		fmt.Println("Setting up network...")
		// Add your network setup logic here
		// fmt.Fprintf(os.Stdout, `{"cniVersion": "%s", "interfaces": [{"name": "eth0"}]}`, config.CniVersion)
	case "DEL":
		fmt.Println("Cleaning up network...")
		// Add your network cleanup logic here
	default:
		fmt.Fprintf(os.Stderr, "Unsupported CNI_COMMAND: %s\n", cmd)
		os.Exit(1)
	}
}
