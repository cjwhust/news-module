package company

import (
	"github.com/astaxie/beego/orm"
)

type News struct {
	Id int `orm:"column(id);auto"`
	Title string `orm:"size(50);unique"`
	Description string `orm:"size(255)"`
	ImgUrl string `orm:"size(255)"`
	Content string `orm:"type(text)"`
	PublishTime string `orm:"size(20)"`
	Type int8
}

func (n *News) TableName() string{
	return "tb_news"
}

func (n *News) Insert() error {
	if _, err := orm.NewOrm().Insert(n); err != nil{
		return err
	}
	return  nil
}

func (n *News) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(n)
}

func GetOneNews(id int) (n *News, err error){
	n = &News{Id: id}
	if err = orm.NewOrm().Read(n); err == nil{
		return n, nil
	}
	return nil, err
}

func FindNewsPagination(id int)(before, after int){
	var newses []News
	news := new(News)
	orm.NewOrm().QueryTable(news).Filter("id__gt", id).OrderBy("id").All(&newses)
	if len(newses) > 0 {
		after = newses[0].Id
	}else {
		after = 0
	}
	newses = nil
	orm.NewOrm().QueryTable(news).Filter("id__lt", id).OrderBy("-id").All(&newses)
	if len(newses) > 0 {
		before = newses[0].Id
	}else {
		before = 0
	}
	return before, after
}

func (n *News) QueryByCondition(fields ...int)(newses []News, total int64, err error){

	qs := n.Query()
	if fields[0] != 0{
		qs = qs.Filter("type__exact", fields[0])
	}
	if fields[1] == 0 {
		total, err = qs.Count()
	}else {
		total, err = qs.Limit(fields[1], fields[2]).OrderBy("-PublishTime").All(&newses)
	}
	if err == nil{
		return newses, total, nil
	}
	return nil, 0, err
}

func (n *News) Delete() error{
	if _, err := orm.NewOrm().Delete(n); err!= nil {
		return  err
	}
	return nil
}

func (n *News) QueryAll() (news []News, total int64, err error){
	qs := n.Query()
	if total, err = qs.All(&news); err == nil {
		return news, total, nil
	}
	return nil, 0, err
}

func (n *News) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(n, fields...); err!=nil{
		return  err
	}
	return nil
}

func (n *News) Read(fields ...string) error{
	if err := orm.NewOrm().Read(n, fields...); err != nil{
		return err
	}
	return nil
}