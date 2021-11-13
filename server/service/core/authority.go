package core

import (
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
)

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)


func (a *AuthorityService) CreateContextLink(contextInfo request.CreateContextLinkReq)(err error,contextLink string) {
        return err,contextLink
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /core/deleteAuthority [post]
func (a *AuthorityService) GetContextByLink(link request.GetContextByLinkReq )(err error,info response.ContextInfo){
  return err,info
}




