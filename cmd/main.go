package main

import (
	"log"
	"log/syslog"
	"os"
)

func main() {

	// Log to syslog
	logWriter, err := syslog.New(syslog.LOG_SYSLOG, "My Awesome App")
	if err != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}
	// set the log output
	log.SetOutput(logWriter)

	// Parse command (ADD/DEL)
	cmd := os.Getenv("CNI_COMMAND")
	args := os.Getenv("CNI_ARGS")

	log.Println("CNI ARGS", args)

	switch cmd {
	case "ADD":
		log.Println("Setting up network...")
		// Add your network setup logic here
		// fmt.Fprintf(os.Stdout, `{"cniVersion": "%s", "interfaces": [{"name": "eth0"}]}`, config.CniVersion)
	case "DEL":
		log.Println("Cleaning up network...")
		// Add your network cleanup logic here
	default:
		log.Printf("Unsupported CNI_COMMAND: %s\n", cmd)
		os.Exit(1)
	}
}
