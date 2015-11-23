package controllers

import (
	"japi/models"
	//"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about AI
type AIController struct {
	beego.Controller
}

// @Title Get
// @Description ai chat
// @Param	chat		path 	string	true		"Please input chat"
// @Success 200 {ai} models.AI
// @Failure 403 :chat is empty
// @router /:chat [get]
func (o *AIController) Say() {
	chat := o.GetString(":chat")
	
	var ai models.AI
	ai.Say(chat)
	
	o.Data["json"] = ai
	o.ServeJson()
}

