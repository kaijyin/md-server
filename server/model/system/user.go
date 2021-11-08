package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type User struct {
	global.GVA_MODEL
	Username string    `json:"userName" gorm:"unique;not null"`
	Password string    `json:"password"  gorm:"not null"`
}
