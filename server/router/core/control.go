package core

import (
	"github.com/gin-gonic/gin"
)

type ControlRouter struct {
}

func (s *ControlRouter) InitSystemRouter(Router *gin.RouterGroup) {
	//sysRouter := Router.Group("control").Use(middleware.OperationRecord())
}
