package company

import "github.com/astaxie/beego/orm"

type Message struct {
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

func (m *Message) TableName() string{
	return "tb_message"
}

func (m *Message) Query() orm.QuerySeter{
	return orm.NewOrm().QueryTable(m)
}

type MessageFlag struct {
	Id int `orm:"column(id);pk"`
	Name string `orm:"size(10)"`
	TypeId int8 `orm:column(type_id)`
}

func (m *MessageFlag) TableName() string{
	return "tb_message_flag"
}

func (m *MessageFlag) Query() orm.QuerySeter{
	return  orm.NewOrm().QueryTable(m)
}

func (m *Message) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return  nil
}

func GetOneMessage(id string)(m *Message, err error){
	m = &Message{Id: id}
	if err = orm.NewOrm().Read(m); err == nil{
		return m, nil
	}
	return nil, err
}

func (m *MessageFlag) QueryAll(typeId int)(mf []MessageFlag, err error){
	qs := m.Query()
	if _,err = qs.Filter("type_id__exact", typeId).All(&mf); err!=nil{
		return
	}
	return
}