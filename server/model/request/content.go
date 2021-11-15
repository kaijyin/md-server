package request

type UserId struct {
	UserId uint `json:"userId"`
}
type CreateCatalogReq struct {
	UserId
	FatherCatalogId uint   `json:"fatherCatalogId"`
	CatalogName     string `json:"catalogName"`
}

type CreateDocumentReq struct {
	UserId
	FatherCatalogId uint   `json:"fatherCatalogId"`
	DocumentName    string `json:"documentName"`
	Content         string `json:"content"`
}

type DeleteCatalogReq struct {
	UserId
	CatalogId uint `json:"catalogId"`
}
type DeleteDocumentReq struct {
	UserId
	DocumentId uint `json:"documentId"`
}

type GetContentByIdReq struct {
	UserId
	DocumentId uint `json:"documentId"`
}

type UpdateContextNameReq struct {
	UserId
	ContextId uint   `json:"contextId"`
	NewName   string `json:"newName"`
}

type UpdateContentReq struct {
	UserId
	DocumentId uint   `json:"documentId"`
	NewContent string `json:"newContent"`
}

// 名称模糊查询+分页+是否按照名称排序

type PageList struct {
	PageInfo
	Desc bool `json:"desc"` // 排序方式:升序false(默认)|降序true
}

type GetContextsInfoReq struct {
	UserId
	FatherCatalogId uint `json:"fatherCatalogId"`
	PageList
}

type GetCatalogsInfoByNameReq struct {
	UserId
	CatalogName string `json:"catalogName"`
	PageList
}
