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

type GoldShopApi struct {
}

var goldShopService = service.ServiceGroupApp.OrderManagerServiceGroup.GoldShopService


// CreateGoldShop 创建goldShop表
// @Tags GoldShop
// @Summary 创建goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldShop true "创建goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldShop/createGoldShop [post]
func (goldShopApi *GoldShopApi) CreateGoldShop(c *gin.Context) {
	var goldShop orderManager.GoldShop
	err := c.ShouldBindJSON(&goldShop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldShopService.CreateGoldShop(&goldShop); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoldShop 删除goldShop表
// @Tags GoldShop
// @Summary 删除goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldShop true "删除goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldShop/deleteGoldShop [delete]
func (goldShopApi *GoldShopApi) DeleteGoldShop(c *gin.Context) {
	id := c.Query("ID")
	if err := goldShopService.DeleteGoldShop(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoldShopByIds 批量删除goldShop表
// @Tags GoldShop
// @Summary 批量删除goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goldShop/deleteGoldShopByIds [delete]
func (goldShopApi *GoldShopApi) DeleteGoldShopByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := goldShopService.DeleteGoldShopByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoldShop 更新goldShop表
// @Tags GoldShop
// @Summary 更新goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orderManager.GoldShop true "更新goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldShop/updateGoldShop [put]
func (goldShopApi *GoldShopApi) UpdateGoldShop(c *gin.Context) {
	var goldShop orderManager.GoldShop
	err := c.ShouldBindJSON(&goldShop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goldShopService.UpdateGoldShop(goldShop); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoldShop 用id查询goldShop表
// @Tags GoldShop
// @Summary 用id查询goldShop表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderManager.GoldShop true "用id查询goldShop表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldShop/findGoldShop [get]
func (goldShopApi *GoldShopApi) FindGoldShop(c *gin.Context) {
	id := c.Query("ID")
	if regoldShop, err := goldShopService.GetGoldShop(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoldShop": regoldShop}, c)
	}
}

// GetGoldShopList 分页获取goldShop表列表
// @Tags GoldShop
// @Summary 分页获取goldShop表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderManagerReq.GoldShopSearch true "分页获取goldShop表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldShop/getGoldShopList [get]
func (goldShopApi *GoldShopApi) GetGoldShopList(c *gin.Context) {
	var pageInfo orderManagerReq.GoldShopSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goldShopService.GetGoldShopInfoList(pageInfo); err != nil {
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
