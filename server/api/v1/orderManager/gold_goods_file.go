package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
    orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoldGoodsFileApi struct {
}

var goldGoodsFileService = service.ServiceGroupApp.OrderManagerServiceGroup.GoldGoodsFileService


// CreateGoldGoodsFile 创建goldGoodsFile表
// @Tags GoldGoodsFile
// @Summary 创建goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldGoodsFile true "创建goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldGoodsFile/createGoldGoodsFile [post]
func (goldGoodsFileApi *GoldGoodsFileApi) CreateGoldGoodsFile(c *gin.Context) {
	var goldGoodsFile orderManager.GoldGoodsFile
	err := c.ShouldBindJSON(&goldGoodsFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldGoodsFileService.CreateGoldGoodsFile(&goldGoodsFile); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoldGoodsFile 删除goldGoodsFile表
// @Tags GoldGoodsFile
// @Summary 删除goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldGoodsFile true "删除goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldGoodsFile/deleteGoldGoodsFile [delete]
func (goldGoodsFileApi *GoldGoodsFileApi) DeleteGoldGoodsFile(c *gin.Context) {
	id := c.Query("ID")
	if err := goldGoodsFileService.DeleteGoldGoodsFile(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoldGoodsFileByIds 批量删除goldGoodsFile表
// @Tags GoldGoodsFile
// @Summary 批量删除goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goldGoodsFile/deleteGoldGoodsFileByIds [delete]
func (goldGoodsFileApi *GoldGoodsFileApi) DeleteGoldGoodsFileByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := goldGoodsFileService.DeleteGoldGoodsFileByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoldGoodsFile 更新goldGoodsFile表
// @Tags GoldGoodsFile
// @Summary 更新goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldGoodsFile true "更新goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldGoodsFile/updateGoldGoodsFile [put]
func (goldGoodsFileApi *GoldGoodsFileApi) UpdateGoldGoodsFile(c *gin.Context) {
	var goldGoodsFile orderManager.GoldGoodsFile
	err := c.ShouldBindJSON(&goldGoodsFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldGoodsFileService.UpdateGoldGoodsFile(goldGoodsFile); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoldGoodsFile 用id查询goldGoodsFile表
// @Tags GoldGoodsFile
// @Summary 用id查询goldGoodsFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderManager.GoldGoodsFile true "用id查询goldGoodsFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldGoodsFile/findGoldGoodsFile [get]
func (goldGoodsFileApi *GoldGoodsFileApi) FindGoldGoodsFile(c *gin.Context) {
	id := c.Query("ID")
	if regoldGoodsFile, err := goldGoodsFileService.GetGoldGoodsFile(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoldGoodsFile": regoldGoodsFile}, c)
	}
}

// GetGoldGoodsFileList 分页获取goldGoodsFile表列表
// @Tags GoldGoodsFile
// @Summary 分页获取goldGoodsFile表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderManagerReq.GoldGoodsFileSearch true "分页获取goldGoodsFile表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldGoodsFile/getGoldGoodsFileList [get]
func (goldGoodsFileApi *GoldGoodsFileApi) GetGoldGoodsFileList(c *gin.Context) {
	var pageInfo orderManagerReq.GoldGoodsFileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goldGoodsFileService.GetGoldGoodsFileInfoList(pageInfo); err != nil {
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
