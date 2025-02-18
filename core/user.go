package core

import (
	"fmt"

	"github.com/a-h/templ"
)

const UserPath string = "/user"

type User struct {
	ID   uint
	Name string `gorm:"unique"`
}

func (u *User) IsReady() bool {
	return u.Name != ""
}

func (u *User) Url() templ.SafeURL {
	return templ.SafeURL(UserUrl(u.ID))
}

func (u *User) Save() error {
	return DB.Save(&u).Error
}

func GetUser(id uint) (*User, error) {
	var user *User
	err := DB.Limit(1).Find(&user, id).Error
	if user == nil {
		user = &User{ID: id}
	}
	return user, err
}

func UserUrl(id uint) string {
	return fmt.Sprintf("%s/%d", UserPath, id)
}
