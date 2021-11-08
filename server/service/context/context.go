package context
import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/model/context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/control"
	"google.golang.org/genproto/protobuf/api"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type ContextService struct {
}

var ContextServiceApp = new(ContextService)

func (apiService *ContextService) CreateContext(api context.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&context.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ContextService) DeleteContext(api context.SysApi) (err error) {
	err = global.GVA_DB.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ContextService) GetContextByNameLick(userId request.GetById, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&context.Context{})

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ContextService) GetAllContexts(userid request.GetById) (err error, contexts []context.Context) {
	err = global.GVA_DB.Table("control").
		Select("context.id context.content").
		Joins("left join context on control.context_id = context.id").
		Where("control.user_id = ?",userid.ID).Scan(&contexts).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (apiService *ContextService) GetContextById(contextId request.GetById) (err error, contxt context.Context) {
	err = global.GVA_DB.Where("id = ?", contextId.ID).First(&contxt).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ContextService) UpdateContextById(contxt context.Context) (err error) {
	var oldC context.Context
	err = global.GVA_DB.Where("id = ?", contxt.ID).First(&oldC).Error
	if err != nil {
		return err
	} else {
		err = global.GVA_DB.Save(&contxt).Error
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ContextService) DeleteContextsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]context.Context{}, "id in ?", ids.Ids).Error
	return err
}
