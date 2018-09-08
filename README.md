# trickster2grafana

Reads origins from a [Trickster](https://github.com/Comcast/trickster) configuration file and POSTs respective datasources to Grafana.

Useful when you have a [multi-origin](https://github.com/Comcast/trickster/blob/master/docs/multi-origin.mdP) Trickster configuration so you don't have to create several Grafana datasources manually.

Only supports the HTTP Pathing multi-origin strategy at this time.

## Example
```
trickster2grafana -grafana-key <your-api-key> -grafana-url http://grafana:3000 -trickster-conf /etc/trickster/trickster.conf -trickster-url http://trickster:9090
```