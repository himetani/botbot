package main

import (
	"flag"
	"log"
	"os"

	"github.com/nlopes/slack"
)

const (
	EnvSlackToken = "SLACK_TOKEN"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", os.Getenv(EnvSlackToken), "")
	flag.StringVar(&token, "token", os.Getenv(EnvSlackToken), "")
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func main() {
	flag.Parse()

	api := slack.New(token)
	os.Exit(run(api))
}
