package system

import "github.com/flipped-aurora/gin-vue-admin/server/router/content"

type RouterGroup struct {
	content.ApiRouter
	BaseRouter
	CasbinRouter
	OperationRecordRouter
	SysRouter
	UserRouter
}
