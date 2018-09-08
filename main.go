package main

const (
	applicationName    = "trickster2grafana"
	applicationVersion = "0.0.1"
)

func main() {
	// Load config from flags
	config := new(config)
	config.loadFromFlags()

	// Load Trickster config
	tricksterConfig := new(tricksterConfig)
	err := tricksterConfig.loadFile(config.tricksterConf)
	if err != nil {
		panic(err)
	}

	// Get resolved values
	resolvedValues, err := resolveHTTPPathing(config, tricksterConfig)
	if err != nil {
		panic(err)
	}

	// Send to Grafana
	err = grafanaPostDatasources(config, resolvedValues)
	if err != nil {
		panic(err)
	}
}
