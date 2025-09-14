package api

type ApiGroup struct {
	VideoContentAnalysisApi
	AiModelApi
	AiChatSessionApi
	AiChatMessageApi
}

var ApiGroupApp = new(ApiGroup)