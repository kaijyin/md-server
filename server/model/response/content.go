package response

type CreateContextResp struct {
	ContextId  uint `json:"contextId"`
}


type GetContextContentResp struct {
	Content string `json:"content"`
}

type CatalogInfo struct {
	ContextId uint `json:"contextId"`
	ContextName uint `json:"contextName"`
}


type CatalogInfoList struct {
	List []CatalogInfo
}

type ContextInfo struct {
	ContextId uint `json:"catalogId"`
	IsCatalog bool `json:"isCatalog"`
	ContextName string `json:"catalogName"`
}
type ContextInfoList struct {
	List []ContextInfo
}

