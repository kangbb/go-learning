package entities

import(
	"time"
)
// UserInfo .
type UserInfo struct {
	UID        int64 `xorm:"autoincr pk 'id'"`//语义标签
	UserName   string
	DepartName string
	CreateAt   *time.Time
}
// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName should not null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}