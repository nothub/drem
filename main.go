package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"hub.lol/drem/version"
	"log"
	"os"
)

var ErrInvalidArgCount = errors.New("invalid arg count")

func main() {
	log.SetFlags(0)

	// see: https://cli.urfave.org/v2/examples/full-api-example/

	app := &cli.App{
		Name:           "drem",
		Version:        version.String(),
		Description:    "Docker rootless environment management",
		Copyright:      "Copyright (c) Florian Hübner",
		DefaultCommand: "help",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list envs",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 0, 0) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "restart",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "logs",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "validate",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 1) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
			{
				Name:  "runas",
				Usage: "",
				Action: func(ctx *cli.Context) error {
					log.Println("TODO")
					return nil
				},
				Before: func(ctx *cli.Context) error {
					if !checkBounds(ctx.NArg(), 1, 255) {
						return ErrInvalidArgCount
					}
					return nil
				},
			},
		},
		Flags:   []cli.Flag{},
		Suggest: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func checkBounds(n int, min uint8, max uint8) (ok bool) {
	if n < int(min) {
		return false
	}
	if n > int(max) {
		return false
	}
	return true
}
