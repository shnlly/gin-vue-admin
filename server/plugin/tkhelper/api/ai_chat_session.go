package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/tkhelper/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AiChatSessionApi struct{}

var aiChatSessionService = service.ServiceGroup{}.AiChatSessionService

// CreateAiChatSession 创建AI对话会话
func (aiChatSessionApi *AiChatSessionApi) CreateAiChatSession(c *gin.Context) {
	var aiChatSession model.AiChatSession
	err := c.ShouldBindJSON(&aiChatSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置用户ID
	aiChatSession.UserId = utils.GetUserID(c)

	// 获取模型信息并设置模型名称
	var aiModel model.AiModel
	if err := global.GVA_DB.Where("id = ?", aiChatSession.ModelId).First(&aiModel).Error; err == nil {
		aiChatSession.ModelName = aiModel.Name
	}

	err = aiChatSessionService.CreateAiChatSession(&aiChatSession)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(aiChatSession, c)
}

// DeleteAiChatSession 删除AI对话会话
func (aiChatSessionApi *AiChatSessionApi) DeleteAiChatSession(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := aiChatSessionService.DeleteAiChatSession(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAiChatSessionByIds 批量删除AI对话会话
func (aiChatSessionApi *AiChatSessionApi) DeleteAiChatSessionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := aiChatSessionService.DeleteAiChatSessionByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAiChatSession 更新AI对话会话
func (aiChatSessionApi *AiChatSessionApi) UpdateAiChatSession(c *gin.Context) {
	var aiChatSession model.AiChatSession
	err := c.ShouldBindJSON(&aiChatSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = aiChatSessionService.UpdateAiChatSession(aiChatSession)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAiChatSession 用id查询AI对话会话
func (aiChatSessionApi *AiChatSessionApi) FindAiChatSession(c *gin.Context) {
	ID := c.Query("ID")
	reaiChatSession, err := aiChatSessionService.GetAiChatSession(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reaiChatSession, c)
}

// GetAiChatSessionList 分页获取AI对话会话列表
func (aiChatSessionApi *AiChatSessionApi) GetAiChatSessionList(c *gin.Context) {
	var pageInfo request.AiChatSessionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 限制只能查看自己的会话
	pageInfo.UserId = utils.GetUserID(c)

	list, total, err := aiChatSessionService.GetAiChatSessionInfoList(pageInfo)
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

// GetUserChatSessions 获取用户的对话会话列表
func (aiChatSessionApi *AiChatSessionApi) GetUserChatSessions(c *gin.Context) {
	userID := utils.GetUserID(c)
	list, err := aiChatSessionService.GetUserChatSessions(userID)
	if err != nil {
		global.GVA_LOG.Error("获取用户会话失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}