package request

type CreateCatalogReq struct {
	FatherCatalogId uint   `json:"fatherCatalogId"`
	CatalogName     string `json:"catalogName"`
}

type CreateDocumentReq struct {
	FatherCatalogId uint   `json:"fatherCatalogId"`
	DocumentName    string `json:"documentName"`
	Content         string `json:"content"`
}

type DeleteDocumentReq struct {
	DocumentId uint `json:"documentId"`
}

type DeleteCatalogReq struct {
	CatalogId uint `json:"catalogId"`
}

type GetContentByIdReq struct {
	DocumentId uint `json:"documentId"`
}

type UpdateContextNameReq struct {
	ContextId uint   `json:"contextId"`
	NewName   string `json:"newName"`
}

type UpdateContentReq struct {
	DocumentId uint   `json:"documentId"`
	NewContent string `json:"newContent"`
}

// 名称模糊查询+分页+是否按照名称排序

type PageList struct {
	PageInfo
	Desc bool `json:"desc"` // 排序方式:升序false(默认)|降序true
}

type GetContextsInfoReq struct {
	FatherCatalogId uint `json:"fatherCatalogId"`
	PageList
}

type GetCatalogsInfoByNameReq struct {
	CatalogName string `json:"catalogName"`
	PageList
}
