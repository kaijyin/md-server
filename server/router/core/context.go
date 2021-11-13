package core

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kaijyin/md-server/server/api/v1"
	"github.com/kaijyin/md-server/server/middleware"
)

type ContextRouter struct {
}

func (s *ContextRouter) InitContextRouter(Router *gin.RouterGroup) {
	contextRouter := Router.Group("context")
	contextRouterApi := v1.ApiGroupApp.CoreApiGroup.ContextApi
	CatalogRouter :=contextRouter.Group("catalog")
	DocumentRouter :=contextRouter.Group("document")
	{
		contextRouter.GET("/:fatherContextId",contextRouterApi.GetContextsInfo)
	}

	{
		CatalogRouter.GET("/:catalogName/:page/:pageSize",contextRouterApi.GetCatalogsInfoByName)
		CatalogRouter.GET("/:catalogName/:page/:pageSize/:desc",contextRouterApi.GetCatalogsInfoByName)

		CatalogRouter.PUT("/:fatherContextId/:catalogName",contextRouterApi.CreateCatalog)
		CatalogRouter.DELETE("/:fatherContextId/:catalogId",contextRouterApi.DeleteContext)
	}

	{
     contextRouter.POST("name",contextRouterApi.CreateContext)
     contextRouter.PUT("name",contextRouterApi.UpdateContextName)
     contextRouter.GET("name",contextRouterApi.GetCatalogsInfoByName)
	}
}
