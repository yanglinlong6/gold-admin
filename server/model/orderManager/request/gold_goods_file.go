package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type GoldGoodsFileSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type AddGoldGoodsFileRequest struct {
	ID         uint       `json:"ID"`                           // 主键ID
	GoodsId    *int       `json:"goodsId" form:"goodsId"`       //商品Id
	FileName   string     `json:"fileName" form:"fileName"`     //文件名
	FileType   *bool      `json:"fileType" form:"fileType"`     //文件类型
	FilePath   string     `json:"filePath" form:"filePath"`     //文件路径
	CreateId   *int       `json:"createId" form:"createId"`     //创建人id
	CreateTime *time.Time `json:"createTime" form:"createTime"` //创建时间
	UpdateId   *int       `json:"updateId" form:"updateId"`     //更新人Id
	UpdateTime *time.Time `json:"updateTime" form:"updateTime"` //更新时间
}
