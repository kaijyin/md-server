package core

import (
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/model/table"
)



type ContextService struct {
}


var ContextServiceApp = new(ContextService)

//先检查是否重命名,在权限控制表和文档表中添加

func (apiService *ContextService) CreateCatalog(req request.CreateCatalogReq) (err error,resp response.CreateContextResp) {


}
func (apiService *ContextService) CreateDocument(req request.CreateDocumentReq) (err error,resp response.CreateContextResp) {

}



//先检查有没有权限,再删除

func (apiService *ContextService) DeleteDocument(req request.DeleteDocumentReq) (err error) {

}

func (apiService *ContextService) DeleteCatalog(req request.DeleteCatalogReq) (err error) {

}



//直接搜索

func (apiService *ContextService) GetCatalogsByName(req request.GetCatalogsInfoByNameReq) (err error,resp response.CatalogInfoList) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var contextList []table.Context
	db:=global.MD_DB.Table("core").
		Select("core.id core.core").
		Joins("left join core on core.context_id = core.id").
		Where("core.user_id = ? and is_catalog = ? and context_name like %?% ", req.UserId,true, req.CatalogName).
		Offset(offset).Limit(limit)
	if req.Desc{
		db=db.Order("context_name desc")
	}else{
		db=db.Order("context_name")
	}
	err=db.Scan(&contextList).Error
	return err, resp
}



//获取用户所有

func (apiService *ContextService) GetContexts(searchReq request.GetContextsInfoReq) (err error, list response.ContextInfoList) {
	limit := searchReq.PageSize
	offset := searchReq.PageSize * (searchReq.Page - 1)
	db:=global.MD_DB.Table("core").
		Select("core.id core.core").
		Joins("left join core on core.context_id = core.id").
		Where("core.user_id = ? and father_catalog_id = ?",searchReq.UserId,searchReq.FatherCatalogId).
		Offset(offset).Limit(limit)
	err=db.Count(&list.Total).Error
	if err !=nil{
		return err, list
	}
	db=db.Order("is_catalog")
	if searchReq.Desc{
		db=db.Order("context_name desc")
	}else{
		db=db.Order("context_name")
	}
	err=db.Scan(&list.ContextsInfo).Error
	return err, list
}


//先检查用户有没有查看权限,有权限再获取

func (apiService *ContextService) GetContentById(req request.GetContentByIdReq) (err error, content response.GetContextContentResp) {
	return err,content
}


//先检查用户有没有权限修改,有权限再改

func (apiService *ContextService) UpdateDocumentContent(req request.UpdateContentReq) (err error) {

	return err
}
func (apiService *ContextService) UpdateContextName(req request.UpdateContextNameReq) (err error) {

	return err
}
