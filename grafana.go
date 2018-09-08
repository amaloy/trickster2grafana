package main

import (
	"fmt"

	"github.com/grafana-tools/sdk"
)

func grafanaPostDatasources(c *config, resolvedValues []*resolved) error {
	fmt.Println()
	fmt.Printf("Going to create these Grafana datasources on %s :", c.grafanaURL)
	fmt.Println()
	fmt.Println()
	for _, r := range resolvedValues {
		fmt.Printf("  NAME: %s   URL: %s", r.id, r.url)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Is this OK? Press enter to continue or Ctrl-C to abort...")
	fmt.Scanln()

	grafana := sdk.NewClient(c.grafanaURL, c.grafanaKey, sdk.DefaultHTTPClient)
	existingDatasources, err := grafana.GetAllDatasources()
	if err != nil {
		return err
	}

	for _, r := range resolvedValues {
		fmt.Printf("%s... ", r.id)

		exists := false
		for _, ds := range existingDatasources {
			if ds.Name == r.id {
				exists = true
				break
			}
		}
		if exists {
			fmt.Print("Already exists, skipping")
		} else {
			var newDS sdk.Datasource
			newDS.Name = r.id
			newDS.Type = "prometheus"
			newDS.URL = r.url
			newDS.Access = "proxy"
			msg, err := grafana.CreateDatasource(newDS)
			if err != nil {
				return err
			}
			fmt.Printf(*msg.Message)
		}
		fmt.Println()
	}

	fmt.Println("Done!")
	return nil
}
