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

type AiChatSessionApi struct{}

var aiChatSessionService = exampleService.AiChatSessionServiceApp

// CreateAiChatSession 创建AI对话会话
// @Tags AiChatSession
// @Summary 创建AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiChatSession true "创建AI对话会话"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /aiChatSession/createAiChatSession [post]
func (aiChatSessionApi *AiChatSessionApi) CreateAiChatSession(c *gin.Context) {
	var aiChatSession example.AiChatSession
	err := c.ShouldBindJSON(&aiChatSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置用户ID
	aiChatSession.UserId = utils.GetUserID(c)

	// 获取模型信息并设置模型名称
	var aiModel example.AiModel
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
// @Tags AiChatSession
// @Summary 删除AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiChatSession true "删除AI对话会话"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /aiChatSession/deleteAiChatSession [delete]
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
// @Tags AiChatSession
// @Summary 批量删除AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /aiChatSession/deleteAiChatSessionByIds [delete]
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
// @Tags AiChatSession
// @Summary 更新AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiChatSession true "更新AI对话会话"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /aiChatSession/updateAiChatSession [put]
func (aiChatSessionApi *AiChatSessionApi) UpdateAiChatSession(c *gin.Context) {
	var aiChatSession example.AiChatSession
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
// @Tags AiChatSession
// @Summary 用id查询AI对话会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query example.AiChatSession true "用id查询AI对话会话"
// @Success 200 {object} response.Response{data=example.AiChatSession,msg=string} "查询成功"
// @Router /aiChatSession/findAiChatSession [get]
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
// @Tags AiChatSession
// @Summary 分页获取AI对话会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query exampleReq.AiChatSessionSearch true "分页获取AI对话会话列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /aiChatSession/getAiChatSessionList [get]
func (aiChatSessionApi *AiChatSessionApi) GetAiChatSessionList(c *gin.Context) {
	var pageInfo exampleReq.AiChatSessionSearch
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
// @Tags AiChatSession
// @Summary 获取用户的对话会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]example.AiChatSession,msg=string} "获取成功"
// @Router /aiChatSession/getUserSessions [get]
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
