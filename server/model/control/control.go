package control

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type Control struct {
	global.GVA_MODEL
	UserId          uint   `json:"userId" gorm:"not null"`           // 角色ID
	Level           uint   `json:"level" gorm:"not null"`            // 目录/文章所在层级
	ContextId       uint   `json:"contextId" gorm:"not null"`        // 目录/文章Id
	FatherContextId uint   `json:"fatherContextId" gorm:"not null;"` // 父目录Id,用于查找和递归删除
	Authority       uint   `json:"control" gorm:"not null;"`       // 用户对文档所有权限
	ext             string `json:"ext"`                              // 扩展字段
	Context         context.Context                                  // 外键引用Context表
	User            system.User                                      // 外键引用User表
}
//```
//json
//{
//"UserId":"所属用户Id",
//"contextId":"文档/目录id",
//"level":"所在层级",
//"flag":"目录文档标志位",
//"menuname":"文档/目录名称",
//"fatherId":"int",//父目录id,根目录为0,用于递归删除
//"permision":"权限int",//三种权限,owner/1,read/2,write/3
//"ext":"扩展字段",
//}
//```