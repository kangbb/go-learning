package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

var engine *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(192.168.99.100:3308)/test?charset=utf8&parseTime=true")
	checkErr(err)

    //sync the struct
    err = engine.Sync2(new(UserInfo))
    checkErr(err)

	//set regular of the table and field name
	engine.SetMapper(core.GonicMapper{})
	engine.ShowSQL(true)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}