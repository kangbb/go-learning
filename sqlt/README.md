# sqlt

一个简单的golang database/sql CRUD操作模板。<br/>
目前已经完成的功能：
```
1.连接数据库
2.创建Table
3.向Table插入数据
4.查询数据(一条或多条)
```
### 使用方式
- **创建数据库连接**
```
energine, err := sqlt.NewEnergine("mysql", "root:root@tcp(192.168.99.100:3308)/test?charset=utf8&parseTime=true")
if err != nil {
    panic(err)
}
```
数据库连接需要根据运行环境自行配置。
- **创建Table或注册Table**
```
type UserInfo struct {
	UID        int64 `myorm:"primary key 'id'"`//语义标签
	UserName   string
	DepartName string
	CreateAt   *time.Time
}
//表格不存在则创建
//表格存在，注册该表格，以便后期使用
db, err := energine.RegisterTable(new(UserInfo), "userinfo")
if err != nil{
    panic(err)
}
```
✦myorm:语义标签，字段名使用`''`括起来;其余标签与mysql列属性相同

- **向Table插入数据**
```
err = db.Save(user)
if err != nil{
    panic(err)
}
```
- **查询数据**
```
//查询语句支持SQL原生语句
//lists类型为slice

//查询所有数据
lists, err :=db.Find(reflect.TypeOf(new(UserInfo)), "SELECT * FROM userinfo")
if err != nil{
    panic(err)
}
fmt.Println(lists)

//查询一条数据
list, err :=db.FindOne(reflect.TypeOf(new(UserInfo)), "SELECT * FROM userinfo")
if err != nil{
    panic(err)
}
fmt.Println(list)
```