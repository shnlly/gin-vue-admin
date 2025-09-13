package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
	AttachmentCategoryApi
	VideoContentAnalysisApi
	AiModelApi
	AiChatSessionApi
	AiChatMessageApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	attachmentCategoryService    = service.ServiceGroupApp.ExampleServiceGroup.AttachmentCategoryService
)
