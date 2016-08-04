package controllers

import (
	"github.com/astaxie/beego"
	"github.com/uuid"
	"strings"
	"os"
	"io"
)

var(
	upload bool
)
type BaseController struct {
	beego.Controller
}

type Meta struct {
	Code int `json:"code"`
	Message string `json:"msg"`
}

type RowsList struct {
	Total int64 `json:"total"`
	Rows interface{} `json:"rows"`
}

type BaseResult struct {
	MetaInfo interface{} `json:"meta"`
	DataInfo interface{} `json:"data"`
}

type QueryModel struct {
	Query, Fields, Order, SortBy string
	PageSize, Page int64

}

func ResultMsg(code int, message string, data interface{}) BaseResult{
	var(
		br = BaseResult{}
		meta = Meta{}
	)
	meta.Code = code
	meta.Message = message
	br.MetaInfo = &meta
	br.DataInfo = &data
	return br
}

func (this *BaseController) SaveFile(para, prefix string)(url, fileName string){
	file, h, err := this.GetFile(para)
	if file == nil {
		return
	}
	if err != nil {
		beego.Error(err)
	}
	fileName = h.Filename
	typename := strings.Split(fileName, ".")
	keyName := prefix +"_"+ uuid.Rand().Raw()+"."+typename[len(typename)-1]
	url = "/static/upload/"+keyName
	f, _ := os.Create("static/upload/"+keyName)
	defer f.Close()
	io.Copy(f, file)
	if upload {
		if flag, name := UploadFile("static/upload/"+keyName, keyName); flag == true{
			keyName = name
			os.Remove("static/upload/"+keyName)
			url = QiniuPrefix+keyName
		}
	}
	return
}

func (this *BaseController) Error(){
	this.Data["json"] = ResultMsg(400, "请求错误", nil)
	this.ServeJSON()
	this.StopRun()
}

func init(){
	upload = beego.AppConfig.DefaultBool("qiniu_enable", false)
}

func (this *BaseController) MakeHeader(){
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET,POST")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
}
