package controllers

import (
	"github.com/qiniu/api.v7/conf"
	"github.com/qiniu/api.v7/kodo"
	"github.com/qiniu/api.v7/kodocli"
	"github.com/astaxie/beego"
)

var (
	QiniuKey = beego.AppConfig.DefaultString("qiniu_key", "")
	QiniuSecret = beego.AppConfig.DefaultString("qiniu_secret", "")
	QiniuPrefix = beego.AppConfig.String("qiniu_prefix")
	QiniuBucket = beego.AppConfig.String("qiniu_bucket")
	Err error
)

func UploadFile(filepath , key string) (bool, string){

	conf.ACCESS_KEY = QiniuKey
	conf.SECRET_KEY = QiniuSecret

	zone := 0
	c := kodo.New(zone, nil)

	policy := &kodo.PutPolicy{
		Scope: QiniuBucket,
		Expires: 3600,
	}

	token := c.MakeUptoken(policy)
	uploader := kodocli.NewUploader(zone, nil)

	var ret kodo.PutRet
	if key == ""{
		Err = uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	}else {
		Err = uploader.PutFile(nil, &ret, token, key, filepath, nil)
	}
	beego.BeeLogger.Info("upload info: ", ret)
	if Err != nil {
		beego.BeeLogger.Error("upload File failed: ", Err)
		return false, ""
	}
	return true, ret.Key
}