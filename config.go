package main

import (
	"flag"
	"os"
)

// Contains application configuration
type config struct {
	tricksterConf string
	tricksterURL  string
	grafanaURL    string
	grafanaKey    string
}

func (c *config) loadFromFlags() {
	f := flag.NewFlagSet(applicationName, flag.ExitOnError)
	f.StringVar(&c.tricksterConf, "trickster-conf", "", "Path to the Trickster conf file")
	f.StringVar(&c.tricksterURL, "trickster-url", "", "Base URL for accessing Trickster")
	f.StringVar(&c.grafanaURL, "grafana-url", "", "Grafana API URL")
	f.StringVar(&c.grafanaKey, "grafana-key", "", "Grafana API key")
	f.Parse(os.Args[1:])

	// Validate
	if len(c.tricksterConf) == 0 ||
		len(c.tricksterURL) == 0 ||
		len(c.grafanaURL) == 0 ||
		len(c.grafanaKey) == 0 {
		f.Usage()
		os.Exit(2)
	}
}
