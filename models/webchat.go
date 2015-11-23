package models

import (
	"fmt"
)

type WebChatTextReceiver struct {
	ToUserName   string
	FromUserName   string
	CreateTime   string
	MsgType   string
	Content   string
	MsgId   string
}

func (rc *WebChatTextReceiver) Return(ai AI) string {
	var str = 	"<xml>" + 
				"<ToUserName>" + 
				rc.FromUserName +
				"</ToUserName>" +
				"<FromUserName>" +
				rc.ToUserName +
				"</FromUserName>" +
				"<CreateTime>" +
				rc.CreateTime +
				"</CreateTime>" +
				"<MsgType>" +
				rc.MsgType +
				"</MsgType>" +
				"<Content>" +
				ai.Chat + 
				"</Content>" +
				"</xml>"
	fmt.Println(str)
	return str
}