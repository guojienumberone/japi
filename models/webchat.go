package models

import (
	//"fmt"
	"encoding/xml"
	"time"
)

type WebChatTextReceiver struct {
	ToUserName   string
	FromUserName   string
	CreateTime   string
	MsgType   string
	Content   string
	MsgId   string
}

type WebChatTextReturner struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName   string
	CreateTime   time.Duration
	MsgType   string
	Content   string
}

func (rc *WebChatTextReceiver) Return(ai AI) WebChatTextReturner {
	var returner WebChatTextReturner
	returner.ToUserName = rc.FromUserName
	returner.FromUserName = rc.ToUserName
	returner.CreateTime = time.Duration(time.Now().Unix())
	returner.MsgType = rc.MsgType
	returner.Content = ai.Chat
	
	return returner
}