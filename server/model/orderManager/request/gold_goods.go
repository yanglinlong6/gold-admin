package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type GoldGoodsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type AddGoldGoodsRequest struct {
	ID                *uint                     `json:"ID"`                // 主键ID
	GoodsTypeId       *int                      `json:"goodsTypeId"`       //商品分类Id
	GoodsName         string                    `json:"goodsName"`         //商品名称
	GoodsPrice        *int                      `json:"goodsPrice"`        //商品价格
	GoldGoodsFileList []AddGoldGoodsFileRequest `json:"goldGoodsFileList"` // 上传文件路径
}
