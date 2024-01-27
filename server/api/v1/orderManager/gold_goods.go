package orderManager

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
	orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoldGoodsApi struct {
}

var goldGoodsService = service.ServiceGroupApp.OrderManagerServiceGroup.GoldGoodsService

// CreateGoldGoods 创建goldGoods表
// @Tags GoldGoods
// @Summary 创建goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldGoods true "创建goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldGoods/createGoldGoods [post]
func (goldGoodsApi *GoldGoodsApi) CreateGoldGoods(c *gin.Context) {
	var goldGoods orderManager.GoldGoods
	err := c.ShouldBindJSON(&goldGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldGoodsService.CreateGoldGoods(&goldGoods); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// AddGoldGoodsAndFiles 创建goldGoods表还有图片文件
// @Tags GoldGoods
// @Summary 创建goldGoods表还有图片文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManagerReq.AddGoldGoodsRequest true "创建goldGoods表还有图片文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldGoods/AddGoldGoodsAndFiles [post]
func (goldGoodsApi *GoldGoodsApi) AddGoldGoodsAndFiles(c *gin.Context) {
	userId, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var addGoldGoodsAndFiles orderManagerReq.AddGoldGoodsRequest
	err = c.ShouldBindJSON(&addGoldGoodsAndFiles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldGoodsService.AddGoldGoodsAndFiles(&addGoldGoodsAndFiles, userId); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoldGoods 删除goldGoods表
// @Tags GoldGoods
// @Summary 删除goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldGoods true "删除goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldGoods/deleteGoldGoods [delete]
func (goldGoodsApi *GoldGoodsApi) DeleteGoldGoods(c *gin.Context) {
	id := c.Query("ID")
	if err := goldGoodsService.DeleteGoldGoods(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoldGoodsByIds 批量删除goldGoods表
// @Tags GoldGoods
// @Summary 批量删除goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goldGoods/deleteGoldGoodsByIds [delete]
func (goldGoodsApi *GoldGoodsApi) DeleteGoldGoodsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := goldGoodsService.DeleteGoldGoodsByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoldGoods 更新goldGoods表
// @Tags GoldGoods
// @Summary 更新goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldGoods true "更新goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldGoods/updateGoldGoods [put]
func (goldGoodsApi *GoldGoodsApi) UpdateGoldGoods(c *gin.Context) {
	var goldGoods orderManager.GoldGoods
	err := c.ShouldBindJSON(&goldGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldGoodsService.UpdateGoldGoods(goldGoods); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoldGoods 用id查询goldGoods表
// @Tags GoldGoods
// @Summary 用id查询goldGoods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderManager.GoldGoods true "用id查询goldGoods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldGoods/findGoldGoods [get]
func (goldGoodsApi *GoldGoodsApi) FindGoldGoods(c *gin.Context) {
	id := c.Query("ID")
	if regoldGoods, err := goldGoodsService.GetGoldGoods(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoldGoods": regoldGoods}, c)
	}
}

// GetGoldGoodsList 分页获取goldGoods表列表
// @Tags GoldGoods
// @Summary 分页获取goldGoods表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderManagerReq.GoldGoodsSearch true "分页获取goldGoods表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldGoods/getGoldGoodsList [get]
func (goldGoodsApi *GoldGoodsApi) GetGoldGoodsList(c *gin.Context) {
	var pageInfo orderManagerReq.GoldGoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goldGoodsService.GetGoldGoodsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
