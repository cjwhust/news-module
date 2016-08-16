package controllers

import (
	m "news-module/models"
	"html/template"
	"time"
	"github.com/uuid"
)

type MainController struct {
	BaseController
}

type MessResponse struct {
	Title string
	Content template.HTML
	PublishTime string
	Description string
	MessFlagName string
}

func (c *MainController) Get() {
	c.TplName = "editor.tpl"
}

func (c *MainController) Content(){
	c.TplName = "content.tpl"
}

func (c *MainController) GetOne(){
	id := c.GetString("id")
	mess, err := m.GetOneMessage(id)
	if err != nil{
		c.TplName = "error.tpl"
	}else {
		var mr MessResponse
		mr.Content = template.HTML(mess.Context)
		mr.PublishTime = mess.PublishTime
		mr.Title = mess.Name
		mr.Description = mess.Description
		mf, _ := m.GetOneMessageFlag(mess.MessageFlagId)
		mr.MessFlagName = mf.Name
		c.Data["Mess"] = mr
		c.TplName = "news.tpl"
	}
}

func (c *MainController) Save(){
	var message m.Message
	message.Name = c.GetString("title")
	message.Time = time.Now().Format("2006-01-02 15:04:05")
	message.Context = c.GetString("content")
	message.MessageFlagId,_ = c.GetInt("flag")
	message.Description = c.GetString("description")
	message.DeleteFlag = "0"
	message.Status = "0"
	message.Id = uuid.Rand().Raw()
	url, _ := c.SaveFile("Image", "news")
	message.ImgUrl = url
	message.MessageTypeId, _ = c.GetInt("type")
	if err :=message.Insert(); err != nil {
		c.Data["json"] = ResultMsg(400, "新增失败", err)
	}else {
		c.Data["json"]= ResultMsg(200, "新增成功", nil)
	}
	c.ServeJSON()
}

func (c *MainController) GetFlags(){
	var mf m.MessageFlag
	typeId, _ := c.GetInt("type")
	l,_ :=mf.QueryAll(typeId)
	c.Data["json"]= ResultMsg(200, "查询成功", l)
	c.ServeJSON()
}
