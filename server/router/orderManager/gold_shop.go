package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoldShopRouter struct {
}

// InitGoldShopRouter 初始化 goldShop表 路由信息
func (s *GoldShopRouter) InitGoldShopRouter(Router *gin.RouterGroup) {
	goldShopRouter := Router.Group("goldShop").Use(middleware.OperationRecord())
	goldShopRouterWithoutRecord := Router.Group("goldShop")
	var goldShopApi = v1.ApiGroupApp.OrderManagerApiGroup.GoldShopApi
	{
		goldShopRouter.POST("createGoldShop", goldShopApi.CreateGoldShop)   // 新建goldShop表
		goldShopRouter.DELETE("deleteGoldShop", goldShopApi.DeleteGoldShop) // 删除goldShop表
		goldShopRouter.DELETE("deleteGoldShopByIds", goldShopApi.DeleteGoldShopByIds) // 批量删除goldShop表
		goldShopRouter.PUT("updateGoldShop", goldShopApi.UpdateGoldShop)    // 更新goldShop表
	}
	{
		goldShopRouterWithoutRecord.GET("findGoldShop", goldShopApi.FindGoldShop)        // 根据ID获取goldShop表
		goldShopRouterWithoutRecord.GET("getGoldShopList", goldShopApi.GetGoldShopList)  // 获取goldShop表列表
	}
}
