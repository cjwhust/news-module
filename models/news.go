package company

import "github.com/astaxie/beego/orm"

type News struct {
	Id string `orm:"column(id);size(32);pk"`
	UserId string `orm:"size(32)"`
	MessageTypeId int
	Name string `orm:"size(200)"`
	PublishTime string `orm:"size(30)"`
	Time string `orm:"size(30)"`
	ParentId string `orm:"size(32)"`
	Status string `orm:"size(1)"`
	DeleteFlag string `orm:"size(1)"`
	Context string `orm:"type(text)"`
	MessageFlagId int
	Description string `orm:"size(200)"`
	ImgUrl string `orm:"size(200)"`
}

func (m *News) TableName() string{
	return "tb_message"
}

func (m *News) Query() orm.QuerySeter{
	return orm.NewOrm().QueryTable(m)
}

type NewsFlag struct {
	Id int `orm:"column(id);pk"`
	Name string `orm:"size(10)"`
	TypeId int8 `orm:column(type_id)`
}

func (m *NewsFlag) TableName() string{
	return "tb_message_flag"
}

func (m *NewsFlag) Query() orm.QuerySeter{
	return  orm.NewOrm().QueryTable(m)
}

func (m *News) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return  nil
}

func (m *News) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err!=nil{
		return  err
	}
	return nil
}

func GetOneMessage(id string)(m *News, err error){
	m = &News{Id: id}
	if err = orm.NewOrm().Read(m); err == nil{
		return m, nil
	}
	return nil, err
}

func GetOneMessageFlag(id int)(m *NewsFlag, err error){
	m = &NewsFlag{Id: id}
	err = orm.NewOrm().Read(m)
	return
}

func (m *NewsFlag) QueryAll(typeId int)(mf []NewsFlag, err error){
	qs := m.Query()
	if _,err = qs.Filter("type_id__exact", typeId).All(&mf); err!=nil{
		return
	}
	return
}