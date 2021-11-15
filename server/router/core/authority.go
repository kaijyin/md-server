package core

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kaijyin/md-server/server/api/v1"
	"github.com/kaijyin/md-server/server/middleware"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.JWTAuth())
	authorityRouterApi := v1.ApiGroupApp.CoreApiGroup.AuthorityApi
	{
		authorityRouter.GET("/:contextCode",authorityRouterApi.GetContextByLink)
		authorityRouter.POST("/:contextId/:permission",authorityRouterApi.CreateContextLink)
	}
}
