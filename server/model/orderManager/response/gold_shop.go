package response

import (
	"time"
)

// goldShop表 结构体  GoldShop
type GoldShopResp struct {
	ID         uint       `json:"ID"`         // 主键ID
	GoodsId    *int       `json:"goodsId"`    //商品Id
	GoodsName  string     `json:"goodsName"`  //商品名称
	BuyId      *int       `json:"buyId"`      //购买人Id
	BuyName    string     `json:"buyName"`    //购买人名称
	Status     *int       `json:"status"`     //购买人Id
	CreateName string     `json:"createName"` //创建人名称
	CreateTime *time.Time `json:"createTime"` //创建时间
	UpdateName string     `json:"updateName"` //更新人名称
	UpdateTime *time.Time `json:"updateTime"` //更新时间
}
