package core

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/zon/hxcore"
)

const UserPath string = "/user"

type User struct {
	ID    uint
	Name  string `gorm:"unique"`
	Ready bool
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
	if err != nil {
		return user, err
	}
	if user == nil || user.ID == 0 {
		user = &User{
			ID:   id,
			Name: hxcore.RandomString(16),
		}
		err = user.Save()
	}
	return user, err
}

func UserUrl(id uint) string {
	return fmt.Sprintf("%s/%d", UserPath, id)
}
