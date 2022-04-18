package model

import (
	"errors"
	"go_template/pkg/database"
	"go_template/pkg/model/common"
	"go_template/pkg/util/encrypt"

	uuid "github.com/satori/go.uuid"
)

var (
	AdminCanNotDelete = "ADMIN_CAN_NOT_DELETE"
	LdapCanNotUpdate  = "LDAP_CAN_NOT_UPDATE"
)

const (
	EN string = "en-US"
	ZH string = "zh-CN"
)

type User struct {
	common.BaseModel
	ID       string `json:"-" gorm:"type:varchar(64)"`
	Name     string `json:"name" gorm:"type:varchar(256);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(256)"`
	Email    string `json:"email" gorm:"type:varchar(256);not null;unique"`
	Language string `json:"language" gorm:"type:varchar(64)"`
	IsAdmin  bool   `json:"-" gorm:"type:boolean;default:false"`
	IsSuper  bool   `json:"-" gorm:"type:boolean;default:false"`
	IsActive bool   `json:"-" gorm:"type:boolean;default:true"`
	Type     string `json:"type" gorm:"type:varchar(64)"`
}

type Token struct {
	Token string `json:"access_token"`
}

func (u *User) BeforeCreate() (err error) {
	u.ID = uuid.NewV4().String()
	return err
}

func (u *User) BeforeDelete() (err error) {
	if u.Name == "admin" {
		return errors.New(AdminCanNotDelete)
	}

	err = database.DB.Model(UserMessage{}).Where("user_id =?", u.ID).Delete(&UserMessage{}).Error
	if err != nil {
		return err
	}
	return err
}

func (u *User) ValidateOldPassword(password string) (bool, error) {
	oldPassword, err := encrypt.StringDecrypt(u.Password)
	if err != nil {
		return false, err
	}
	if oldPassword != password {
		return false, err
	}
	return true, err
}
