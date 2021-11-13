package utils

var (

CreateContextLinkVerify = Rules{ "DocumentId": {NotEmpty()}, "Permission": {NotEmpty()}}
GetContextByLinkVerify  = Rules{ "ContextLink": {NotEmpty()}}


	CreateCatalogVerify     = Rules{"UserId":{NotEmpty()},"FatherCatalogId":{NotEmpty()},"CatalogName": {NotEmpty()}}
	CreateDocumentVerify     = Rules{"UserId":{NotEmpty()},"FatherCatalogId":{NotEmpty()},"DocumentName": {NotEmpty()},
		"Content":{NotEmpty()}}

	CatalogIdVerify                = Rules{"CatalogId":{NotEmpty()}}
	DocumentIdVerify                = Rules{"DocumentId":{NotEmpty()}}

	UpdateContextNameVerify = Rules{"ContextId":{NotEmpty()},"NewName":{NotEmpty()}}
	UpdateDocumentContentVerify =Rules{"DocumentId":{NotEmpty()},"NewContent":{NotEmpty()}}

	GetCatalogsInfoByNameVerify = Rules{"CatalogName": {NotEmpty()},"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	GetContextsInfoVerify = Rules{"UserId":{NotEmpty()},"FatherCatalogId":{NotEmpty()},"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}

	LoginVerify             = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	ChangePasswordVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}

)
