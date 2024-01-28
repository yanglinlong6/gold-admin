package response

import (
	"time"
)

// goldGoods表 结构体  GoldGoods
type GoldGoodsResp struct {
	ID                uint                `json:"ID"`                // 主键ID
	GoodsTypeId       *int                `json:"goodsTypeId"`       //商品分类Id
	GoodsName         string              `json:"goodsName"`         //商品名称
	GoodsPrice        *int                `json:"goodsPrice"`        //商品价格
	CreateName        string              `json:"createName"`        //商品创建人
	CreateTime        *time.Time          `json:"createTime"`        //创建时间
	UpdateName        string              `json:"updateName"`        //商品更新人
	UpdateTime        *time.Time          `json:"updateTime"`        //更新时间
	GoldGoodsFileList []GoldGoodsFileResp `json:"goldGoodsFileList"` //商品文件列表
}

// goldGoodsFile表 结构体  GoldGoodsFile
type GoldGoodsFileResp struct {
	ID         uint       `json:"ID"`         // 主键ID
	GoodsId    *uint      `json:"goodsId"`    //商品Id
	FileName   string     `json:"fileName"`   //文件名
	FileType   *bool      `json:"fileType"`   //文件类型
	FilePath   string     `json:"filePath"`   //文件路径
	CreateName string     `json:"createName"` //创建人创建人
	CreateTime *time.Time `json:"createTime"` //创建时间
	UpdateName string     `json:"updateName"` //更新人创建人
	UpdateTime *time.Time `json:"updateTime"` //更新时间
}
