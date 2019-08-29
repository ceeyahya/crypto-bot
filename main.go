package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ceeyahya/crypto-bot/utils"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(
		"xoxb-730854609473-739104563735-0zwrWCG6HGlWlXNj9Jc770LC",
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "crypto-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					utils.Answer(rtm, ev, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}

}
