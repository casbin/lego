package cmd

import (
	"github.com/casbin/lego/v4/log"
	"github.com/urfave/cli"
)

func Before(ctx *cli.Context) error {
	if ctx.GlobalString("path") == "" {
		log.Fatal("Could not determine current working directory. Please pass --path.")
	}

	err := createNonExistingFolder(ctx.GlobalString("path"))
	if err != nil {
		log.Fatalf("Could not check/create path: %v", err)
	}

	if ctx.GlobalString("server") == "" {
		log.Fatal("Could not determine current working server. Please pass --server.")
	}

	return nil
}
