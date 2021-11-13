package table

import (
	"github.com/kaijyin/md-server/server/global"
)

type User struct {
	global.GVA_MODEL
	Username string `json:"userName" gorm:"index;unique;not null"`
	Password string `json:"password"  gorm:"not null"`
	Email    string `json:"email"  gorm:"not null"`
}
