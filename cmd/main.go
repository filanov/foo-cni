package main

import (
	"log"
	"os"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	cni "github.com/containernetworking/cni/pkg/version"
)

var cniVersion = "0.3.1"

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
	}, cni.All, cniVersion)
}

func addCmd(args *skel.CmdArgs) error {
	log.Printf("Add iface: %s args: %s\n", args.IfName, args.Args)
	log.Printf("Add config: %s", string(args.StdinData))

	result := &current.Result{CNIVersion: current.ImplementedSpecVersion}
	return types.PrintResult(result, cniVersion)
}

func delCmd(args *skel.CmdArgs) error {
	log.Printf("del iface: %s args: %s\n", args.IfName, args.Args)
	log.Printf("del config: %s", string(args.StdinData))
	result := &current.Result{CNIVersion: current.ImplementedSpecVersion}
	return types.PrintResult(result, cniVersion)
}

func checkCmd(args *skel.CmdArgs) error {
	log.Printf("Check %+v", args)
	result := &current.Result{CNIVersion: current.ImplementedSpecVersion}
	return types.PrintResult(result, cniVersion)
}

func gcCmd(args *skel.CmdArgs) error {
	log.Printf("GC %+v", args)
	result := &current.Result{CNIVersion: current.ImplementedSpecVersion}
	return types.PrintResult(result, cniVersion)
}

func statusCmd(args *skel.CmdArgs) error {
	log.Printf("Status %+v", args)
	result := &current.Result{CNIVersion: current.ImplementedSpecVersion}
	return types.PrintResult(result, cniVersion)
}
