package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"hub.lol/drem/buildinfo"

	_ "github.com/urfave/cli/v2"
)

type Cmd struct {
	name    string
	desc    string
	minArgs uint8
	maxArgs uint8
	run     func(args []string) uint8
}

func (cmd *Cmd) Run(args []string) uint8 {
	if len(args) < int(cmd.minArgs) {
		log.Fatalln("invalid argument count ( < min )")
	}
	if len(args) > int(cmd.maxArgs) {
		log.Fatalln("invalid argument count ( > max )")
	}
	return cmd.run(args)
}

var cmds = make(map[string]Cmd)

func main() {
	log.SetFlags(0)

	var help bool
	flag.BoolVar(&help, "help", false, "Print message help and exit")

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")

	flag.Parse()

	var cmd string
	if len(flag.Args()) > 0 {
		cmd = flag.Args()[0]
	}
	cmd = strings.TrimSpace(cmd)
	cmd = strings.ToLower(cmd)

	var args []string
	if len(flag.Args()) > 1 {
		args = flag.Args()[1:]
	}

	log.Printf("cmd=%q args=%q\n", cmd, args)

	cmds["list"] = Cmd{
		name:    "list",
		desc:    "",
		minArgs: 0,
		maxArgs: 0,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["create"] = Cmd{
		name:    "create",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["delete"] = Cmd{
		name:    "delete",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["start"] = Cmd{
		name:    "start",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["stop"] = Cmd{
		name:    "stop",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["restart"] = Cmd{
		name:    "restart",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["logs"] = Cmd{
		name:    "logs",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["status"] = Cmd{
		name:    "status",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["validate"] = Cmd{
		name:    "validate",
		desc:    "",
		minArgs: 1,
		maxArgs: 1,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["runas"] = Cmd{
		name:    "runas",
		desc:    "",
		minArgs: 1,
		maxArgs: 255,
		run:     func([]string) uint8 { return 0 },
	}

	cmds["help"] = Cmd{
		name:    "help",
		desc:    "",
		minArgs: 0,
		maxArgs: 1,
		run: func([]string) uint8 {
			printHelp()
			return 0
		},
	}

	cmds["version"] = Cmd{
		name:    "version",
		desc:    "",
		minArgs: 0,
		maxArgs: 0,
		run: func([]string) uint8 {
			printVersion()
			return 0
		},
	}

	// alias root cmd to help cmd
	cmds[""] = cmds["help"]

	c, ok := cmds[cmd]
	// print help if cmd is unknown
	if !ok {
		c = cmds["help"]
		os.Exit(1)
	}

	os.Exit(int(c.Run(args)))
}

func printHelp() {
	fmt.Println("TODO")
}

func printVersion() {
	fmt.Printf("%s %s\n", buildinfo.Name(), buildinfo.Version())
}
