package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(255);comment:用户名"`
	Password string `gorm:"type:char(32);comment:密码"`
	Status   int8   `gorm:"type:tinyint(1);default:0;comment: 0正常 1冻结"`
}
