package utils

import (
	"strings"

	"github.com/ceeyahya/crypto-bot/handlers"
	"github.com/nlopes/slack"
)

// Answer Handles messages sent to the bot
func Answer(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	// rLink := "http://api.coinlayer.com/api/live"
	var rsp string
	txt := msg.Text
	txt = strings.TrimPrefix(txt, prefix)
	txt = strings.TrimSpace(txt)

	if len(txt) > 5 {
		rsp = "I don't think that's a Cryptocurrency try something along the lines of 'BTC' or 'ETH'"
		rtm.SendMessage(rtm.NewOutgoingMessage(rsp, msg.Channel))
	} else if txt == "help" {
		rsp = "Hey There, I'm CryptoBot a nice .. Bot. You can tag me on any channel and add a cryptocurrency symbol and I'll get you the prices. Try me, I'm the best at what I do."
		rtm.SendMessage(rtm.NewOutgoingMessage(rsp, msg.Channel))
	} else {
		resp := handlers.GetCurrency(txt)
		rtm.SendMessage(rtm.NewOutgoingMessage(txt+" sells for "+resp+" $", msg.Channel))
	}
}
