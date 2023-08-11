package main

import (
	"fmt"
	"log"
	"os"
)

type cmd struct {
	argsMin uint8
	argsMax uint8
	usage   string
	funk    func(args []string)
}

func (cmd cmd) exec(args []string) {
	if len(args) < int(cmd.argsMin) || len(args) > int(cmd.argsMax) {
		log.Printf("usage: %s\n", cmd.usage)
		os.Exit(2)
	}
	cmd.funk(args)
}

var cmds = make(map[string]cmd)

func init() {

	cmds["list"] = cmd{
		argsMin: 0,
		argsMax: 0,
		usage:   "list",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["create"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "create <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["delete"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "delete <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["start"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "start <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["stop"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "stop <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["restart"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "restart <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["status"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "status <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["logs"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "logs <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["validate"] = cmd{
		argsMin: 1,
		argsMax: 1,
		usage:   "validate <name>",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

	cmds["runas"] = cmd{
		argsMin: 2,
		argsMax: 255,
		usage:   "runas <name> <cmd> [<arg>...]",
		funk: func(args []string) {
			fmt.Println("TODO") // TODO
		}}

}
