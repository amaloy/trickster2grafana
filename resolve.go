package main

import (
	"net/url"
	"path"
)

// resolved contains values that are ready for output (e.g. to Grafana)
type resolved struct {
	id  string
	url string
}

// Resolve for Tricker's HTTP Pathing multi-origin strategy
func resolveHTTPPathing(c *config, tc *tricksterConfig) ([]*resolved, error) {
	all := make([]*resolved, 0)
	for origin := range tc.Origins {
		r := new(resolved)
		r.id = origin

		u, err := url.Parse(c.tricksterURL)
		if err != nil {
			return nil, err
		}
		u.Path = path.Join(u.Path, r.id)
		r.url = u.String()

		all = append(all, r)
	}
	return all, nil
}
