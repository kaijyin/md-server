package core

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kaijyin/md-server/server/api/v1"
)

type ContextRouter struct {
}

func (s *ContextRouter) InitContextRouter(Router *gin.RouterGroup) {
	contextRouter := Router.Group("context")
	contextRouterApi := v1.ApiGroupApp.CoreApiGroup.ContextApi
	catalogRouter :=contextRouter.Group("catalog")
	documentRouter :=contextRouter.Group("document")
	{
		contextRouter.GET("/:fatherCatalogId",contextRouterApi.GetContextsInfo)
	}

	{
		catalogRouter.GET("/:catalogName/:page/:pageSize",contextRouterApi.GetCatalogsInfoByName)
		catalogRouter.GET("/:catalogName/:page/:pageSize/:desc",contextRouterApi.GetCatalogsInfoByName)

		catalogRouter.POST("/:fatherCatalogId/:catalogName",contextRouterApi.CreateCatalog)
		catalogRouter.PUT("/:contextId/:newName",contextRouterApi.UpdateContextName)
		catalogRouter.DELETE("/:catalogId",contextRouterApi.DeleteCatalog)
	}

	{
       documentRouter.GET("/:documentId",contextRouterApi.GetContentById)
       documentRouter.POST("/:fatherCatalogId/:contextName",contextRouterApi.CreateDocument)
       documentRouter.PUT("/:contextId/:newName",contextRouterApi.UpdateContextName)
       documentRouter.DELETE("/:documentId",contextRouterApi.DeleteDocument)
	}
}
