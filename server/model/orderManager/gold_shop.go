// 自动生成模板GoldShop
package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// goldShop表 结构体  GoldShop
type GoldShop struct {
	global.GVA_MODEL
	GoodsId    *int       `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品Id;size:10;"`     //商品Id
	BuyId      *int       `json:"buyId" form:"buyId" gorm:"column:buy_id;comment:购买人Id;size:10;"`          //购买人Id
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:订单状态;size:4;"`          //购买人Id
	CreateId   *int       `json:"createId" form:"createId" gorm:"column:create_id;comment:创建人Id;size:10;"` //创建人Id
	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`    //创建时间
	UpdateId   *int       `json:"updateId" form:"updateId" gorm:"column:update_id;comment:更新人Id;size:10;"` //更新人Id
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`    //更新时间
}

// TableName goldShop表 GoldShop自定义表名 gold_shop
func (GoldShop) TableName() string {
	return "gold_shop"
}
