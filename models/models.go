package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Sex    		int64  `json:"sex"`
}

func RegisterDB() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:xkl1992qq@/cloudworld?charset=utf8")//密码为空格式
}
func Read() []*User {
	//o := orm.NewOrm()
	//user := User{Id: 1}
	//
	//err := o.Read(&user)
	//
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.Id, user.Name,user.Sex,user.Phone)
	//}
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("user")
	// 也可以直接使用对象作为表名
	//user := new(User)
	//qs = o.QueryTable(user) // 返回 QuerySeter
	var users []*User
	num, err := qs.All(&users)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	return users
}
func insert() {
	o := orm.NewOrm()
	var user User
	user.Name = "slene"
	user.Sex = 1

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}