package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoldGoodsRouter struct {
}

// InitGoldGoodsRouter 初始化 goldGoods表 路由信息
func (s *GoldGoodsRouter) InitGoldGoodsRouter(Router *gin.RouterGroup) {
	goldGoodsRouter := Router.Group("goldGoods").Use(middleware.OperationRecord())
	goldGoodsRouterWithoutRecord := Router.Group("goldGoods")
	var goldGoodsApi = v1.ApiGroupApp.OrderManagerApiGroup.GoldGoodsApi
	{
		goldGoodsRouter.POST("createGoldGoods", goldGoodsApi.CreateGoldGoods)             // 新建goldGoods表
		goldGoodsRouter.POST("addGoldGoodsAndFiles", goldGoodsApi.AddGoldGoodsAndFiles)   // 新建goldGoods表还有图片
		goldGoodsRouter.DELETE("deleteGoldGoods", goldGoodsApi.DeleteGoldGoods)           // 删除goldGoods表
		goldGoodsRouter.DELETE("deleteGoldGoodsByIds", goldGoodsApi.DeleteGoldGoodsByIds) // 批量删除goldGoods表
		goldGoodsRouter.PUT("updateGoldGoods", goldGoodsApi.UpdateGoldGoods)              // 更新goldGoods表
	}
	{
		goldGoodsRouterWithoutRecord.GET("findGoldGoods", goldGoodsApi.FindGoldGoods)       // 根据ID获取goldGoods表
		goldGoodsRouterWithoutRecord.GET("getGoldGoodsList", goldGoodsApi.GetGoldGoodsList) // 获取goldGoods表列表
	}
}
