package main

import (
	"github.com/mpaulweeks/slackbot/robots"
	"github.com/mpaulweeks/slackbot/server"
)

type robutt struct {
}

func (r *robutt) Run(p *robots.Payload) string {
	return "hello"
}

func (r *robutt) Description() string {
	return "Create a review request for a random contributor."
}

func (r *robutt) HandleButton(bp *robots.ButtonPayload) string {
	return "success"
}

func main() {
	robots.RegisterRobot("roulette", &robutt{})
	server.Main(robots.Robots)
}
