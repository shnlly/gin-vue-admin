package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	exampleService "github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AiChatMessageApi struct{}

var aiChatMessageService = exampleService.AiChatMessageServiceApp

// CreateAiChatMessage 创建AI对话消息
// @Tags AiChatMessage
// @Summary 创建AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiChatMessage true "创建AI对话消息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /aiChatMessage/createAiChatMessage [post]
func (aiChatMessageApi *AiChatMessageApi) CreateAiChatMessage(c *gin.Context) {
	var aiChatMessage example.AiChatMessage
	err := c.ShouldBindJSON(&aiChatMessage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = aiChatMessageService.CreateAiChatMessage(&aiChatMessage)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAiChatMessage 删除AI对话消息
// @Tags AiChatMessage
// @Summary 删除AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiChatMessage true "删除AI对话消息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /aiChatMessage/deleteAiChatMessage [delete]
func (aiChatMessageApi *AiChatMessageApi) DeleteAiChatMessage(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := aiChatMessageService.DeleteAiChatMessage(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAiChatMessageByIds 批量删除AI对话消息
// @Tags AiChatMessage
// @Summary 批量删除AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /aiChatMessage/deleteAiChatMessageByIds [delete]
func (aiChatMessageApi *AiChatMessageApi) DeleteAiChatMessageByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := aiChatMessageService.DeleteAiChatMessageByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAiChatMessage 更新AI对话消息
// @Tags AiChatMessage
// @Summary 更新AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiChatMessage true "更新AI对话消息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /aiChatMessage/updateAiChatMessage [put]
func (aiChatMessageApi *AiChatMessageApi) UpdateAiChatMessage(c *gin.Context) {
	var aiChatMessage example.AiChatMessage
	err := c.ShouldBindJSON(&aiChatMessage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = aiChatMessageService.UpdateAiChatMessage(aiChatMessage)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAiChatMessage 用id查询AI对话消息
// @Tags AiChatMessage
// @Summary 用id查询AI对话消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query example.AiChatMessage true "用id查询AI对话消息"
// @Success 200 {object} response.Response{data=example.AiChatMessage,msg=string} "查询成功"
// @Router /aiChatMessage/findAiChatMessage [get]
func (aiChatMessageApi *AiChatMessageApi) FindAiChatMessage(c *gin.Context) {
	ID := c.Query("ID")
	reaiChatMessage, err := aiChatMessageService.GetAiChatMessage(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reaiChatMessage, c)
}

// GetAiChatMessageList 分页获取AI对话消息列表
// @Tags AiChatMessage
// @Summary 分页获取AI对话消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query exampleReq.AiChatMessageSearch true "分页获取AI对话消息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /aiChatMessage/getAiChatMessageList [get]
func (aiChatMessageApi *AiChatMessageApi) GetAiChatMessageList(c *gin.Context) {
	var pageInfo exampleReq.AiChatMessageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := aiChatMessageService.GetAiChatMessageInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSessionMessages 获取会话消息列表
// @Tags AiChatMessage
// @Summary 获取会话消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param sessionId query uint true "会话ID"
// @Success 200 {object} response.Response{data=[]example.AiChatMessage,msg=string} "获取成功"
// @Router /aiChatMessage/getSessionMessages [get]
func (aiChatMessageApi *AiChatMessageApi) GetSessionMessages(c *gin.Context) {
	sessionID := c.Query("sessionId")
	if sessionID == "" {
		response.FailWithMessage("会话ID不能为空", c)
		return
	}

	sessionIDUint, parseErr := strconv.ParseUint(sessionID, 10, 32)
	if parseErr != nil {
		response.FailWithMessage("会话ID格式错误", c)
		return
	}

	list, err := aiChatMessageService.GetSessionMessages(uint(sessionIDUint))
	if err != nil {
		global.GVA_LOG.Error("获取会话消息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}

// SendMessage 发送消息
// @Tags AiChatMessage
// @Summary 发送消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body exampleReq.ChatRequest true "发送消息请求"
// @Success 200 {object} response.Response{msg=string} "发送成功"
// @Router /aiChatMessage/sendMessage [post]
func (aiChatMessageApi *AiChatMessageApi) SendMessage(c *gin.Context) {
	var req exampleReq.ChatRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)

	// 记录请求信息
	global.GVA_LOG.Info("收到发送消息请求",
		zap.Uint("session_id", req.SessionId),
		zap.Uint("model_id", req.ModelId),
		zap.Bool("stream", req.Stream),
		zap.Int("message_length", len(req.Message)),
		zap.Uint("user_id", userID),
	)

	if req.Stream {
		// 设置SSE响应头
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Access-Control-Allow-Origin", "*")

		err = aiChatMessageService.SendChatMessage(req, userID, c.Writer)
		if err != nil {
			global.GVA_LOG.Error("发送消息失败!", zap.Error(err))
			c.SSEvent("error", err.Error())
			return
		}
	} else {
		err = aiChatMessageService.SendChatMessage(req, userID, nil)
		if err != nil {
			global.GVA_LOG.Error("发送消息失败!", zap.Error(err))
			response.FailWithMessage("发送失败: "+err.Error(), c)
			return
		}
		response.OkWithMessage("发送成功", c)
	}
}
