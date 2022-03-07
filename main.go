package main

import (
	"log"
	"os"

	"github.com/hazelops/tflock/dynamodb"
	"github.com/urfave/cli/v2"
)

var Version = "deployment"

var (
	lockID  string
	region  string
	profile string
	table   string
)

func main() {
	app := &cli.App{
		Name:  "tflock",
		Usage: "lock terraform state",
		UsageText: `
		# Lock terraform state in S3 (DynamoBD)
		tflock --lock-id nutcorp-tf-state/env/terraform.tfstate
		`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lock-id",
				Usage:       "specify lock id, usually just a path to the state file in your s3 bucket",
				Required:    true,
				Destination: &lockID,
			},
			&cli.StringFlag{
				Name:        "aws-region",
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
			&cli.StringFlag{
				Name:        "dynamodb-table",
				Usage:       "specify DynamoDB table",
				Required:    true,
				Destination: &table,
				Value:       "tf-state-lock",
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
