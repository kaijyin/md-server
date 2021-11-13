package response

type ContextLinkResp struct {
	ContextLink string `json:"contextLink"`
}

type GetContextByLinkReq struct {
	ContextInfo ContextInfo `json:"contextInfo"`
}