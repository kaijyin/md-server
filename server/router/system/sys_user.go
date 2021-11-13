package system

import (

	"github.com/gin-gonic/gin"
	v1 "github.com/kaijyin/md-server/server/api/v1"
	"github.com/kaijyin/md-server/server/middleware"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("", baseApi.ChangePassword)         // 用户修改密码
		userRouter.DELETE("", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("", baseApi.SetUserInfo)                // 设置用户信息
	}
	userRouterWithoutRecord.GET("", baseApi.GetUserInfo)  // 获取自身信息
}
