package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	help := flag.Bool("help", false, "show help message")
	flag.Parse()

	if *help {
		fmt.Println("TODO") // TODO
		os.Exit(0)
	}

	var cmd string
	var args []string
	if len(flag.Args()) > 1 {
		cmd = flag.Args()[1]
	}
	if len(flag.Args()) > 2 {
		args = flag.Args()[2:]
	}

	c, ok := cmds[cmd]
	if !ok {
		log.Println("no such command")
		os.Exit(2)
	}

	c.exec(args)
}

func binPath() string {
	return ""
}

func detox() string {
	return ""
}
