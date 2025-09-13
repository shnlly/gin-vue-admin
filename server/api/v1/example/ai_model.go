package example

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	exampleService "github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AiModelApi struct{}

var aiModelService = exampleService.AiModelServiceApp

// CreateAiModel 创建AI模型
// @Tags AiModel
// @Summary 创建AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiModel true "创建AI模型"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /aiModel/createAiModel [post]
func (aiModelApi *AiModelApi) CreateAiModel(c *gin.Context) {
	var aiModel example.AiModel
	err := c.ShouldBindJSON(&aiModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = aiModelService.CreateAiModel(&aiModel)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAiModel 删除AI模型
// @Tags AiModel
// @Summary 删除AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiModel true "删除AI模型"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /aiModel/deleteAiModel [delete]
func (aiModelApi *AiModelApi) DeleteAiModel(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := aiModelService.DeleteAiModel(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAiModelByIds 批量删除AI模型
// @Tags AiModel
// @Summary 批量删除AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /aiModel/deleteAiModelByIds [delete]
func (aiModelApi *AiModelApi) DeleteAiModelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := aiModelService.DeleteAiModelByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAiModel 更新AI模型
// @Tags AiModel
// @Summary 更新AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.AiModel true "更新AI模型"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /aiModel/updateAiModel [put]
func (aiModelApi *AiModelApi) UpdateAiModel(c *gin.Context) {
	var aiModel example.AiModel
	err := c.ShouldBindJSON(&aiModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = aiModelService.UpdateAiModel(aiModel)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAiModel 用id查询AI模型
// @Tags AiModel
// @Summary 用id查询AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query example.AiModel true "用id查询AI模型"
// @Success 200 {object} response.Response{data=example.AiModel,msg=string} "查询成功"
// @Router /aiModel/findAiModel [get]
func (aiModelApi *AiModelApi) FindAiModel(c *gin.Context) {
	ID := c.Query("ID")
	reaiModel, err := aiModelService.GetAiModel(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reaiModel, c)
}

// GetAiModelList 分页获取AI模型列表
// @Tags AiModel
// @Summary 分页获取AI模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query exampleReq.AiModelSearch true "分页获取AI模型列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /aiModel/getAiModelList [get]
func (aiModelApi *AiModelApi) GetAiModelList(c *gin.Context) {
	var pageInfo exampleReq.AiModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := aiModelService.GetAiModelInfoList(pageInfo)
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

// GetAiModelPublic 不鉴权的AI模型接口
// @Tags AiModel
// @Summary 不鉴权的AI模型接口
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=example.AiModel,msg=string} "获取成功"
// @Router /aiModel/getAiModelPublic [get]
func (aiModelApi *AiModelApi) GetAiModelPublic(c *gin.Context) {
	aiModelService.GetAiModelPublic()
	response.OkWithMessage("获取成功", c)
}

// SetDefaultAiModel 设置默认AI模型
// @Tags AiModel
// @Summary 设置默认AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path string true "模型ID"
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /aiModel/setDefault [post]
func (aiModelApi *AiModelApi) SetDefaultAiModel(c *gin.Context) {
	idStr := c.Query("ID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("ID格式错误", c)
		return
	}

	err = aiModelService.SetDefaultAiModel(uint(id))
	if err != nil {
		global.GVA_LOG.Error("设置默认模型失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetDefaultAiModel 获取默认AI模型
// @Tags AiModel
// @Summary 获取默认AI模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=example.AiModel,msg=string} "获取成功"
// @Router /aiModel/getDefault [get]
func (aiModelApi *AiModelApi) GetDefaultAiModel(c *gin.Context) {
	aiModel, err := aiModelService.GetDefaultAiModel()
	if err != nil {
		global.GVA_LOG.Error("获取默认模型失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(aiModel, c)
}

// GetEnabledAiModels 获取启用的AI模型列表
// @Tags AiModel
// @Summary 获取启用的AI模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]example.AiModel,msg=string} "获取成功"
// @Router /aiModel/getEnabled [get]
func (aiModelApi *AiModelApi) GetEnabledAiModels(c *gin.Context) {
	list, err := aiModelService.GetEnabledAiModels()
	if err != nil {
		global.GVA_LOG.Error("获取启用模型失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}
