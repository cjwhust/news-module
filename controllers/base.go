package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/cache"
	"strings"
	"qiniupkg.com/x/errors.v7"
	"fmt"
	"strconv"
	"frontsurf/utils"
	"github.com/uuid"
	"os"
	"io"
)

var(
	cpt *captcha.Captcha
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

func (this *BaseController) Display(tpl ...string){
	realURI := strings.Split(this.Ctx.Request.RequestURI,"?")
	this.Data["head_info"] = realURI[0]
	if this.GetSession("userid") != nil {
		this.Data["username"] = this.GetSession("username")
	}else {
		this.Data["username"] = ""
	}
}

func (this *BaseController) GetPage()int{
	var page = 1
	if this.GetString("page") != ""{
		page, _ = strconv.Atoi(this.GetString("page"))
	}
	return page
}

func (this *BaseController) GetLimit()int{
	var limit = utils.LimitAdmin
	if this.GetString("pagesize") != ""{
		limit, _ = strconv.Atoi(this.GetString("pagesize"))
	}
	return limit
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
	url = "/static/base/upload/"+keyName
	f, _ := os.Create("static/base/upload/"+keyName)
	defer f.Close()
	io.Copy(f, file)
	if upload {
		if flag, name := utils.UploadFile("static/base/upload/"+keyName, keyName); flag == true{
			keyName = name
			os.Remove("static/base/upload/"+keyName)
			url = utils.QiniuPrefix+keyName
		}
	}
	return
}

func (this *BaseController) Pagination(total int){
	url := this.Ctx.Request.RequestURI
	page := this.GetPage()
	if this.GetString("tar_info") == "" && this.GetString("page") == ""{
		this.Data["pagination"] = utils.NewPagination(page, total, url+"?page=")
	}else if this.GetString("tar_info") != "" && this.GetString("page") == ""{
		this.Data["pagination"] = utils.NewPagination(page, total, url+"&page=")
	}else {
		this.Data["pagination"] = utils.NewPagination(page, total, url[0 : strings.LastIndex(url, "page=")+5])
	}
}

func (this *BaseController) Error(){
	this.Data["json"] = ResultMsg(400, "请求错误", nil)
	this.ServeJSON()
	this.StopRun()
}

func init(){
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	upload = utils.BeeConf.DefaultBool("qiniu_enable", false)
}

func (this *BaseController) MakeHeader(){
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET,POST")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
}

func (this *BaseController) VerifyCaptcha() bool{
	captcha := this.GetString("captcha")
	captchaId := this.GetString("captcha_id")
	return cpt.Verify(captchaId, captcha)
}

func (this *BaseController) BuildJsonMap(qm QueryModel) (f, s, o []string, q map[string]string, l, of int64){
	var (
		fields, sortBy, order []string
		query map[string]string = make(map[string]string)
		limit, offset int64 = 10, 0
	)
	if v := qm.Fields; v != ""{
		fields = strings.Split(v, ",")
	}
	if v := qm.PageSize; v != 0{
		limit = v
	}
	if v := qm.Page; v != 0{
		offset = v
	}
	if v := qm.SortBy; v != ""{
		sortBy = strings.Split(v, ",")
	}
	if v := qm.Order; v != ""{
		order = strings.Split(v, ",")
	}
	if v := qm.Query; v != ""{
		for _, cond := range strings.Split(v, ","){
			kv := strings.Split(cond, ":")
			if len(kv) != 2{
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
			//utils.Log.Debug(query)
			fmt.Println(query)
		}
	}
	return fields, sortBy, order, query, limit, (offset-1)*limit

}