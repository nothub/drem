package main

import (
	"flag"
	"fmt"
	_ "github.com/samber/lo"
	"github.com/spf13/pflag"
	"hub.lol/drem/buildinfo"
	"log"
)

var cmds = make(map[string]func([]string))

func main() {
	log.SetFlags(0)

	var help bool
	pflag.BoolVarP(&help, "help", "h", false, "Print message help and exit")

	var verbose bool
	pflag.BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	flag.Parse()

	var cmd string
	if len(flag.Args()) > 0 {
		cmd = flag.Args()[0]
	}

	var args []string
	if len(flag.Args()) > 1 {
		args = flag.Args()[1:]
	}

	log.Printf("cmd=%q args=%q\n", cmd, args)

	cmds[""] = func([]string) {
		fmt.Println("Hello, World!")
	}

	cmds["list"] = nil

	cmds["create"] = nil

	cmds["delete"] = nil

	cmds["start"] = nil

	cmds["stop"] = nil

	cmds["restart"] = nil

	cmds["logs"] = nil

	cmds["status"] = nil

	cmds["validate"] = nil

	cmds["runas"] = nil

	cmds["help"] = func([]string) {
		printHelp()
	}

	cmds["version"] = func([]string) {
		printVersion()
	}

	cmds[cmd](args)
}

func printHelp() {
	fmt.Println("TODO")
}

func printVersion() {
	fmt.Printf("%s %s\n", buildinfo.Name(), buildinfo.Version())
}
