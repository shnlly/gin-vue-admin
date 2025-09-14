package router

type RouterGroup struct {
	VideoContentAnalysisRouter
	AiModelRouter
	AiChatSessionRouter
	AiChatMessageRouter
}

var RouterGroupApp = new(RouterGroup)