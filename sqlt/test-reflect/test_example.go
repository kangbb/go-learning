package main

import (
	"github.com/kangbb/go-learning/sqlt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"fmt"
)

type UserInfo struct {
	UID        int64 `myorm:"primary key 'id'"`//语义标签
	UserName   string
	DepartName string
	CreateAt   *time.Time
}
func main()  {
	energine, err := sqlt.NewEnergine("mysql", "root:root@tcp(192.168.99.100:3308)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	t := time.Now()

	user := UserInfo{UID:6, UserName:"kangbb", DepartName:"School", CreateAt:&t}
	db, err := energine.RegisterTable(new(UserInfo), "userinfo")
	if err != nil{
		panic(err)
	}
	err = db.Save(user)
	if err != nil{
		panic(err)
	}

	//lists, err :=db.Find(reflect.TypeOf(new(UserInfo)), "SELECT * FROM userinfo")
	//if err != nil{
	//	panic(err)
	//}
	//for _, v:= range lists{
	//	fmt.Println(v)
	//}

	list, err :=db.FindOne(reflect.TypeOf(new(UserInfo)), "SELECT * FROM userinfo")
	if err != nil{
		panic(err)
	}
	fmt.Println(list)

}