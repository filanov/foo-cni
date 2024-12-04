package main

import (
	"fmt"
	"log"
	"os"

	"github.com/containernetworking/cni/pkg/skel"
	cni "github.com/containernetworking/cni/pkg/version"
)

func main() {
	logFile, err := os.OpenFile("/var/log/foo-cni.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("failed to open log file")
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:    addCmd,
		Del:    delCmd,
		Check:  checkCmd,
		GC:     gcCmd,
		Status: statusCmd,
	}, cni.All, "v1")
}

func addCmd(args *skel.CmdArgs) error {
	log.Printf("Add iface: %s args: %s\n", args.IfName, args.Args)
	log.Printf("Add config: %s", string(args.StdinData))

	fmt.Println("{}")
	return nil
}

func delCmd(args *skel.CmdArgs) error {
	log.Printf("del iface: %s args: %s\n", args.IfName, args.Args)
	log.Printf("del config: %s", string(args.StdinData))
	fmt.Println("{}")
	return nil
}

func checkCmd(args *skel.CmdArgs) error {
	log.Printf("Check %+v", args)
	fmt.Println("{}")
	return nil
}

func gcCmd(args *skel.CmdArgs) error {
	log.Printf("GC %+v", args)
	fmt.Println("{}")
	return nil
}

func statusCmd(args *skel.CmdArgs) error {
	log.Printf("Status %+v", args)
	fmt.Println("{}")
	fmt.Println("{}")
	return nil
}
