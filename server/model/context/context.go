package context

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Context struct {
	global.GVA_MODEL
	ContextType global.ContextType `json:"flag" gorm:"not null"`        // 目录/文章标志位
	ContextName string             `json:"contextName" gorm:"not null"` // 目录/文章名称
	Content     string             `json:"context" gorm:""`                     // 文章内容
}
