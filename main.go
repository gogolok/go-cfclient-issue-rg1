package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudfoundry/go-cfclient/v3/client"
	"github.com/cloudfoundry/go-cfclient/v3/config"
)

func main() {
	cfUser := os.Getenv("CF_USERNAME")
	cfPassword := os.Getenv("CF_PASSWORD")
	cfAPIUrl := os.Getenv("CF_API")

	cfg, err := config.New(cfAPIUrl, config.UserPassword(cfUser, cfPassword))
	if err != nil {
		fmt.Printf("err = %v\n", err)
		os.Exit(1)
	}

	cf, err := client.New(cfg)
	if err != nil {
		fmt.Printf("err = %v\n", err)
		os.Exit(1)
	}

	refreshInterval := 4 * 60 * time.Second // 4 minutes

	ctx := context.Background()
	for {
		stacks, _, err := cf.Stacks.List(ctx, nil)
		if err != nil {
			fmt.Printf("err = %v\n", err)
			//os.Exit(1)
		}
		fmt.Printf("stacks = %v ; err = %v\n", stacks, err)

		fmt.Printf("Sleeping for time interval %v .\n\n", refreshInterval)
		time.Sleep(refreshInterval)
	}
}
