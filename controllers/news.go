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

func (c *MainController) GetInfo(){
	id := c.GetString("id")
	mess, err := m.GetOneMessage(id)
	if err != nil{
		c.Data["json"] = ResultMsg(400, "查询失败", err)
	}else {
		c.Data["json"]= ResultMsg(200, "查询成功", mess)
	}
	c.ServeJSON()
}

func (c *MainController) Save(){
	var news m.News
	if c.GetString("id") == ""{
		news.Name = c.GetString("title")
		news.Time = time.Now().Format("2006-01-02 15:04:05")
		news.Context = c.GetString("content")
		news.MessageFlagId,_ = c.GetInt("flag")
		news.Description = c.GetString("description")
		news.DeleteFlag = "0"
		news.Status = "0"
		news.Id = uuid.Rand().Raw()
		url, _ := c.SaveFile("Image", "news")
		if(url == "") {url="/static/upload/base.png"}
		news.ImgUrl = url
		news.MessageTypeId, _ = c.GetInt("type")
		if err := news.Insert(); err != nil {
			c.Data["json"] = ResultMsg(400, "新增失败", err)
		}else {
			c.Data["json"]= ResultMsg(200, "新增成功", nil)
		}
	}else {
		mess, _ := m.GetOneMessage(c.GetString("id"))
		mess.Name = c.GetString("title")
		mess.Time = time.Now().Format("2006-01-02 15:04:05")
		mess.Context = c.GetString("content")
		mess.MessageFlagId,_ = c.GetInt("flag")
		mess.Description = c.GetString("description")
		mess.DeleteFlag = "0"
		mess.Status = "0"
		if url, _ := c.SaveFile("Image", "scenery"); url != ""{
			mess.ImgUrl = url
		}
		mess.MessageTypeId, _ = c.GetInt("type")
		if err := mess.Update(); err != nil {
			c.Data["json"] = ResultMsg(400, "修改失败", err)
		}else {
			c.Data["json"]= ResultMsg(200, "修改成功", nil)
		}
	}
	c.ServeJSON()
}

func (c *MainController) GetFlags(){
	var mf m.NewsFlag
	typeId, _ := c.GetInt("type")
	l,_ :=mf.QueryAll(typeId)
	c.Data["json"]= ResultMsg(200, "查询成功", l)
	c.ServeJSON()
}
