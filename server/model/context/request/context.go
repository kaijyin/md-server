package request

import (
	"md-server/server/model/common/request"
)

// 名称模糊查询+分页+是否按照名称排序

type SearchContextParams struct {
	UserId uint `json:"userId"`
	ContextName string `json:"contextName"`
	request.PageInfo
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type UpdateContext struct {
	UserId uint `json:"userId"`
	NewContext context.Context
}

