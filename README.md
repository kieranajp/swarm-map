# Swarm-Map

Takes a list of recent checkins on Swarm, removes the last n hours' worth of data (to protect privacy), and plots the remaining dataset on a pretty Mapbox map.

## Installation

Clone this down, run `dep ensure`, copy .env.example to .env, and set your secrets in there. Then `go run main.go` to fire up the server (listens on 8989 by default).
