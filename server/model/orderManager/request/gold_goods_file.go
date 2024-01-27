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
	ID         *uint      `json:"ID"`         // 主键ID
	GoodsId    *int       `json:"goodsId"`    //商品Id
	FileName   string     `json:"fileName"`   //文件名
	FileType   *bool      `json:"fileType"`   //文件类型
	FilePath   string     `json:"filePath"`   //文件路径
	CreateId   *int       `json:"createId"`   //创建人id
	CreateTime *time.Time `json:"createTime"` //创建时间
	UpdateId   *int       `json:"updateId"`   //更新人Id
	UpdateTime *time.Time `json:"updateTime"` //更新时间
}
