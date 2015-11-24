package controllers

import (
	"japi/models"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"net/url"
	//"io/ioutil"
)

// Operations about AI
type WebChatController struct {
	beego.Controller
}

// @Title createUser
// @Description get webchat content
// @Param	body		body 	models.WebChat	true		"body for webchat content"
// @Success 200 {int} models.WebChat.Id
// @Failure 403 body is empty
// @router / [post]
func (wc *WebChatController) Post() {
	fmt.Println("WebChat_Post")
	fmt.Println(string(wc.Ctx.Input.RequestBody))
	
	var receiver models.WebChatTextReceiver
	var returner models.WebChatTextReturner
	
	err:=xml.Unmarshal(wc.Ctx.Input.RequestBody, &receiver)
	if err==nil {
		var ai models.AI
		ai.Say(receiver.Content)	
		
		returner = receiver.Return(ai)
	} else {
		fmt.Println(err)
	}
	
	wc.Data["xml"] = returner
	wc.ServeXml()
}

// @Title Get
// @Description WebChat Check
// @Success 200 :echostr
// @Failure 403 :echostr is empty
// @router / [get]
func (wc *WebChatController) Get() {
	fmt.Println("WebChat_Get")
	fmt.Println(wc.Ctx.Input.Request.URL)
	
	urlParam := wc.Ctx.Input.Request.URL.RawQuery
	m,_ := url.ParseQuery(urlParam)
	
	if (m["echostr"]!=nil){
		wc.Ctx.WriteString(m["echostr"][0])
	} else {
		wc.Ctx.WriteString("No echostr Param")
	}
}
