package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudfoundry/go-cfclient/v3/client"
	"github.com/cloudfoundry/go-cfclient/v3/config"
)

func main() {
	cfUser := os.Getenv("CF_USERNAME")
	cfPassword := os.Getenv("CF_PASSWORD")
	cfAPIUrl := os.Getenv("CF_API")
	cfOrgGuid := os.Getenv("CF_ORG_GUID")

	if len(cfOrgGuid) < 1 {
		fmt.Printf("CF_ORG_GUID must be set!")
		os.Exit(1)
	}

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

	ctx := context.Background()
	org, err := cf.Organizations.Get(ctx, cfOrgGuid)
	if err != nil {
		fmt.Printf("err = %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("org suspended = %v\n", org.Suspended)

	opts := client.NewOrganizationListOptions()
	opts.GUIDs = client.Filter{Values: []string{cfOrgGuid}}

	orgs, _, err := cf.Organizations.List(ctx, opts)
	if err != nil {
		fmt.Printf("err = %v\n", err)
		os.Exit(1)
	}

	for _, org := range orgs {
		fmt.Printf("org name = %v , suspended = %v, guid = %v\n", org.Name, org.Suspended, org.GUID)
	}

}
