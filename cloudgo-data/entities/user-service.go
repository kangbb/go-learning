package entities

import (
	"fmt"
)

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(NewUserInfo(*u))
	checkErr(err)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	everyone := make([]UserInfo, 0)
	err := engine.Find(&everyone)
	if err != nil{
		panic(err)
	}
	fmt.Println(everyone)
	return everyone
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var user  = UserInfo{}
	_, err := engine.Where("id = ?", id).Get(&user)
	if err != nil{
		panic(err)
	}
	return &user
}

