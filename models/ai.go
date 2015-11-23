package models

import (
	"strings"
)

type AI struct {
	Chat   string
}

func (ai *AI) Say(chat string) {
	chat = strings.TrimSpace(chat)
	chat = strings.Replace(chat, " ", "+", -1)
	
	//Language Translation
	var tuling TuLing
	tuling.Translation(chat)
	
	//Return Info
	ai.Chat = tuling.Text;
}

