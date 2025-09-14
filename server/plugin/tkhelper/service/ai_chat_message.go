package service

import (
	"context"
	"fmt"
	"io"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/utils/openai"
	"go.uber.org/zap"
)

type AiChatMessageService struct{}

// CreateAiChatMessage 创建AI对话消息记录
func (aiChatMessageService *AiChatMessageService) CreateAiChatMessage(aiChatMessage *model.AiChatMessage) (err error) {
	err = global.GVA_DB.Create(aiChatMessage).Error
	return err
}

// DeleteAiChatMessage 删除AI对话消息记录
func (aiChatMessageService *AiChatMessageService) DeleteAiChatMessage(ID string, userID uint) (err error) {
	// 软删除AI对话消息记录
	err = global.GVA_DB.Delete(&model.AiChatMessage{}, "id = ?", ID).Error
	return err
}

// DeleteAiChatMessageByIds 批量删除AI对话消息记录
func (aiChatMessageService *AiChatMessageService) DeleteAiChatMessageByIds(IDs []string, deleted_by uint) (err error) {
	// 批量软删除AI对话消息记录
	err = global.GVA_DB.Where("id in ?", IDs).Delete(&model.AiChatMessage{}).Error
	return err
}

// UpdateAiChatMessage 更新AI对话消息记录
func (aiChatMessageService *AiChatMessageService) UpdateAiChatMessage(aiChatMessage model.AiChatMessage) (err error) {
	err = global.GVA_DB.Model(&model.AiChatMessage{}).Where("id = ?", aiChatMessage.ID).Updates(&aiChatMessage).Error
	return err
}

// GetAiChatMessage 根据ID获取AI对话消息记录
func (aiChatMessageService *AiChatMessageService) GetAiChatMessage(ID string) (aiChatMessage model.AiChatMessage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aiChatMessage).Error
	return
}

// GetAiChatMessageInfoList 分页获取AI对话消息记录
func (aiChatMessageService *AiChatMessageService) GetAiChatMessageInfoList(info request.AiChatMessageSearch) (list []model.AiChatMessage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AiChatMessage{})
	var aiChatMessages []model.AiChatMessage

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", *info.StartCreatedAt, *info.EndCreatedAt)
	}
	if info.SessionId != 0 {
		db = db.Where("session_id = ?", info.SessionId)
	}
	if info.Role != "" {
		db = db.Where("role = ?", info.Role)
	}
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Order("message_index ASC")
	err = db.Find(&aiChatMessages).Error
	return aiChatMessages, total, err
}

// GetSessionMessages 获取会话的所有消息
func (aiChatMessageService *AiChatMessageService) GetSessionMessages(sessionID uint) (list []model.AiChatMessage, err error) {
	err = global.GVA_DB.Where("session_id = ?", sessionID).
		Order("message_index ASC").
		Find(&list).Error
	return
}

// SendChatMessage 发送聊天消息并获取AI回复
func (aiChatMessageService *AiChatMessageService) SendChatMessage(req request.ChatRequest, userID uint, responseWriter io.Writer) error {
	// 获取AI模型信息
	var aiModel model.AiModel
	if err := global.GVA_DB.Where("id = ? AND enabled = ?", req.ModelId, true).First(&aiModel).Error; err != nil {
		return fmt.Errorf("AI模型不存在或未启用")
	}

	// 获取会话信息
	var session model.AiChatSession
	if err := global.GVA_DB.Where("id = ? AND user_id = ?", req.SessionId, userID).First(&session).Error; err != nil {
		return fmt.Errorf("会话不存在或无权限访问")
	}

	// 获取会话历史消息
	var historyMessages []model.AiChatMessage
	if err := global.GVA_DB.Where("session_id = ?", req.SessionId).
		Order("message_index ASC").
		Find(&historyMessages).Error; err != nil {
		return fmt.Errorf("获取历史消息失败: %v", err)
	}

	// 保存用户消息
	nextIndex := len(historyMessages)
	userMessage := model.AiChatMessage{
		SessionId:    req.SessionId,
		Role:         "user",
		Content:      req.Message,
		MessageIndex: nextIndex,
	}
	if err := global.GVA_DB.Create(&userMessage).Error; err != nil {
		return fmt.Errorf("保存用户消息失败: %v", err)
	}

	// 创建OpenAI客户端
	client := openai.NewClient(&aiModel)

	// 构建聊天请求
	chatRequest := openai.BuildChatRequest(aiModel.ModelName, historyMessages, req.Message, req.Stream)

	ctx := context.Background()

	if req.Stream {
		// 流式响应
		return aiChatMessageService.handleStreamResponse(ctx, client, chatRequest, req.SessionId, nextIndex+1, responseWriter)
	} else {
		// 非流式响应
		return aiChatMessageService.handleNormalResponse(ctx, client, chatRequest, req.SessionId, nextIndex+1)
	}
}

// handleStreamResponse 处理流式响应
func (aiChatMessageService *AiChatMessageService) handleStreamResponse(ctx context.Context, client *openai.Client, request openai.ChatRequest, sessionID uint, messageIndex int, responseWriter io.Writer) error {
	// 创建自定义流处理器，集成消息保存逻辑
	handler := &ChatStreamHandler{
		BaseHandler: &openai.SSEStreamHandler{Writer: responseWriter},
		service:     aiChatMessageService,
		sessionID:   sessionID,
		messageIndex: messageIndex,
	}

	return client.ChatStream(ctx, request, handler)
}

// handleNormalResponse 处理非流式响应
func (aiChatMessageService *AiChatMessageService) handleNormalResponse(ctx context.Context, client *openai.Client, request openai.ChatRequest, sessionID uint, messageIndex int) error {
	response, err := client.Chat(ctx, request)
	if err != nil {
		return err
	}

	// 保存AI回复消息
	aiMessage := openai.ConvertToChatMessage(sessionID, "assistant", response.Content, messageIndex, response.TokenCount)
	if err := global.GVA_DB.Create(aiMessage).Error; err != nil {
		return fmt.Errorf("保存AI回复消息失败: %v", err)
	}

	// 更新会话信息
	return aiChatMessageService.updateSessionInfo(sessionID, response.Content)
}

// updateSessionInfo 更新会话信息
func (aiChatMessageService *AiChatMessageService) updateSessionInfo(sessionID uint, lastMessage string) error {
	// 截取最后消息的预览
	preview := openai.TruncateMessage(lastMessage, 100)

	return global.GVA_DB.Model(&model.AiChatSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"last_message": preview,
		}).Error
}

// ChatStreamHandler 自定义流处理器，继承SSE处理器并添加数据库保存逻辑
type ChatStreamHandler struct {
	BaseHandler  openai.StreamHandler
	service      *AiChatMessageService
	sessionID    uint
	messageIndex int
}

// OnMessage 处理消息片段
func (h *ChatStreamHandler) OnMessage(content string) error {
	return h.BaseHandler.OnMessage(content)
}

// OnComplete 处理完成事件
func (h *ChatStreamHandler) OnComplete(response openai.ChatResponse) error {
	// 先调用基础处理器
	if err := h.BaseHandler.OnComplete(response); err != nil {
		return err
	}

	// 保存AI回复消息到数据库
	aiMessage := openai.ConvertToChatMessage(h.sessionID, "assistant", response.Content, h.messageIndex, response.TokenCount)
	if err := global.GVA_DB.Create(aiMessage).Error; err != nil {
		global.GVA_LOG.Error("保存AI回复消息失败", zap.Error(err))
		return fmt.Errorf("保存AI回复消息失败: %v", err)
	}

	// 更新会话信息
	if err := h.service.updateSessionInfo(h.sessionID, response.Content); err != nil {
		global.GVA_LOG.Error("更新会话信息失败", zap.Error(err))
	}

	return nil
}

// OnError 处理错误事件
func (h *ChatStreamHandler) OnError(err error) error {
	return h.BaseHandler.OnError(err)
}