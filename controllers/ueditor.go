package controllers

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/astaxie/beego"
	"github.com/uuid"
	"encoding/base64"
	"io/ioutil"
	"regexp"
	"log"
	"os"
)

type UeditorController struct {
	BaseController
}

func (uc *UeditorController) UEController(){
	options := uc.Input().Get("action")
	switch options {
	case "config":
		if file, err := os.Open("conf/config.json"); err !=nil{
			log.Fatal(err)
			os.Exit(1)
		}else {
			defer file.Close()
			fd, _ :=ioutil.ReadAll(file)
			src := string(fd)
			re, _ :=regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/")
			src = re.ReplaceAllString(src, "")
			if(uc.Input().Get("callback") == ""){
				tt := []byte(src)
				var r interface{}
				ffjson.Unmarshal(tt, &r)
				uc.Data["json"] = r
				uc.ServeJSON()
			}else {
				uc.Ctx.WriteString(uc.Input().Get("callback")+"("+src+")")
			}
		}
	case "uploadimage", "uploadfile", "uploadvideo":
		url, fileName := uc.SaveFile("upfile", "ueditor")
		//src, _ := ffjson.Marshal(map[string]interface{}{"state":"SUCCESS","url": url,"title":h.Filename,"original":h.Filename})
		//uc.Ctx.WriteString(string(src))
		uc.Data["json"] = map[string]interface{}{"state":"SUCCESS","url": url,"title": fileName,"original": fileName}
		uc.ServeJSON()
	case "uploadscrawl":
		scrawl := uc.Input().Get("upfile")
		buffer, _ := base64.StdEncoding.DecodeString(scrawl)
		keyName := "ueditor_"+uuid.Rand().Raw()+".jpg"
		url := "/static/base/upload/"+keyName
		if err := ioutil.WriteFile("static/base/upload/"+keyName, buffer, 0666); err != nil {
			beego.BeeLogger.Error("upload scrawl error: ", err)
		}else {
			if flag, name := UploadFile("static/base/upload/"+keyName, keyName); flag == true{
				keyName = name
				os.Remove("static/base/upload/"+keyName)
				url = QiniuPrefix+keyName
			}
		}
		uc.Data["json"] = map[string]interface{}{"state":"SUCCESS","url": url,"title":keyName,"original":keyName}
		uc.ServeJSON()
	}
}