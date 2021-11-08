package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	BaseApi
	SystemApi
}

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
