// 自动生成模板GoldGoods
package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// goldGoods表 结构体  GoldGoods
type GoldGoods struct {
      global.GVA_MODEL
      GoodsTypeId  *int `json:"goodsTypeId" form:"goodsTypeId" gorm:"column:goods_type_id;comment:商品分类Id;size:10;"`  //商品分类Id 
      GoodsName  string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:255;"`  //商品名称 
      GoodsPrice  *int `json:"goodsPrice" form:"goodsPrice" gorm:"column:goods_price;comment:商品价格;size:10;"`  //商品价格 
      CreateId  *int `json:"createId" form:"createId" gorm:"column:create_id;comment:商品创建人;size:10;"`  //商品创建人 
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`  //创建时间 
      UpdateId  *int `json:"updateId" form:"updateId" gorm:"column:update_id;comment:商品更新人;size:10;"`  //商品更新人 
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`  //更新时间 
}


// TableName goldGoods表 GoldGoods自定义表名 gold_goods
func (GoldGoods) TableName() string {
  return "gold_goods"
}

