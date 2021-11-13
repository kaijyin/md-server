package request

// 名称模糊查询+分页+是否按照名称排序

type CreateContextLinkReq struct {
	ContextId uint `json:"contextId"`
	Permission string `json:"permission"`
}

type GetContextByLinkReq struct {
	ContextLink string `json:"contextLink"`
}

