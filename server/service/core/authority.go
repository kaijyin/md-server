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

func (a *AuthorityService) GetContextByLink(link request.GetContextByLinkReq )(err error,info response.ContextInfo){
  return err,info
}




