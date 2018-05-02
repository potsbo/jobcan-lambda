package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/potsbo/jobcan/account"
	"github.com/potsbo/jobcan/config"
)

func jobcanTouch(ctx context.Context, c config.Config) (string, error) {
	if !c.Valid() {
		return "", errors.Errorf("Not enough credential given")
	}
	mode := os.Getenv("JOBCAN_AID")
	if len(mode) == 0 {
		return "", errors.Errorf("env var JOBCAN_AID is empty")
	}

	a := account.New(c)
	// TODO: get errors
	a.Login()
	a.ExecAttendance(mode)
	return "", nil
}

func main() {
	lambda.Start(jobcanTouch)
}
