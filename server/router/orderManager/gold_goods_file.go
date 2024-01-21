package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoldGoodsFileRouter struct {
}

// InitGoldGoodsFileRouter 初始化 goldGoodsFile表 路由信息
func (s *GoldGoodsFileRouter) InitGoldGoodsFileRouter(Router *gin.RouterGroup) {
	goldGoodsFileRouter := Router.Group("goldGoodsFile").Use(middleware.OperationRecord())
	goldGoodsFileRouterWithoutRecord := Router.Group("goldGoodsFile")
	var goldGoodsFileApi = v1.ApiGroupApp.OrderManagerApiGroup.GoldGoodsFileApi
	{
		goldGoodsFileRouter.POST("createGoldGoodsFile", goldGoodsFileApi.CreateGoldGoodsFile)   // 新建goldGoodsFile表
		goldGoodsFileRouter.DELETE("deleteGoldGoodsFile", goldGoodsFileApi.DeleteGoldGoodsFile) // 删除goldGoodsFile表
		goldGoodsFileRouter.DELETE("deleteGoldGoodsFileByIds", goldGoodsFileApi.DeleteGoldGoodsFileByIds) // 批量删除goldGoodsFile表
		goldGoodsFileRouter.PUT("updateGoldGoodsFile", goldGoodsFileApi.UpdateGoldGoodsFile)    // 更新goldGoodsFile表
	}
	{
		goldGoodsFileRouterWithoutRecord.GET("findGoldGoodsFile", goldGoodsFileApi.FindGoldGoodsFile)        // 根据ID获取goldGoodsFile表
		goldGoodsFileRouterWithoutRecord.GET("getGoldGoodsFileList", goldGoodsFileApi.GetGoldGoodsFileList)  // 获取goldGoodsFile表列表
	}
}
