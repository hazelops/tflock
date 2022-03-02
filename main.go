package main

import (
	"log"
	"os"

	"github.com/hazelops/tflock/dynamodb"
	"github.com/urfave/cli/v2"
)

var Version = "deployment"

var lockID string
var region string
var profile string

func main() {
	app := &cli.App{
		Name:  "tflock",
		Usage: "lock terraform state",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lock-id",
				Usage:       "lock id",
				Required:    true,
				Destination: &lockID,
			},
			&cli.StringFlag{
				Name:        "aws region",
				Usage:       "AWS region",
				Required:    true,
				EnvVars:     []string{"TFLOCK_AWS_REGION", "AWS_REGION"},
				Destination: &region,
				Value:       "us-east-1",
			},
			&cli.StringFlag{
				Name:        "aws-profile",
				Usage:       "AWS profile",
				Required:    true,
				EnvVars:     []string{"TFLOCK_AWS_PROFILE", "AWS_PROFILE"},
				Destination: &profile,
				Value:       "default",
			},
		},
		Action: func(c *cli.Context) error {
			dynamodb.Lock(lockID, region, profile)
			return nil
		},
		Version: Version,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
