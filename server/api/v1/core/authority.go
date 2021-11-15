package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/utils"
	"go.uber.org/zap"
)

type AuthorityApi struct {
}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /core/createAuthority [post]
func (a *AuthorityApi) CreateContextLink(c *gin.Context) {
	var link request.CreateContextLinkReq
	_ = c.ShouldBindJSON(&link)
	if err := utils.Verify(link, utils.CreateContextLinkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, contextLink := AuthorityService.CreateContextLink(link); err != nil {
		global.MD_LOG.Error("分享链接创建失败!", zap.Any("err", err))
		response.FailWithMessage("分享链接创建失败"+err.Error(), c)
	} else {
		response.OkWithData(contextLink,c)
	}
}


// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /core/deleteAuthority [post]
func (a *AuthorityApi) GetContextByLink(c *gin.Context) {
	var req request.GetContextByLinkReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.GetContextByLinkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err,contextInfo:= AuthorityService.GetContextByLink(req); err != nil {
		global.MD_LOG.Error("共享文件获取失败!", zap.Any("err", err))
		response.FailWithMessage("共享文件获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(contextInfo, "共享文件获取成功", c)
	}
}
