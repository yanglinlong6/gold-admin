// 自动生成模板GoldGoodsFile
package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// goldGoodsFile表 结构体  GoldGoodsFile
type GoldGoodsFile struct {
      global.GVA_MODEL
      GoodsId  *uint `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品Id;size:10;"`  //商品Id 
      FileName  string `json:"fileName" form:"fileName" gorm:"column:file_name;comment:文件名;size:255;"`  //文件名 
      FileType  *bool `json:"fileType" form:"fileType" gorm:"column:file_type;comment:文件类型;"`  //文件类型 
      FilePath  string `json:"filePath" form:"filePath" gorm:"column:file_path;comment:文件路径;size:1024;"`  //文件路径 
      CreateId  *int `json:"createId" form:"createId" gorm:"column:create_id;comment:创建人id;size:10;"`  //创建人id 
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`  //创建时间 
      UpdateId  *int `json:"updateId" form:"updateId" gorm:"column:update_id;comment:更新人Id;size:10;"`  //更新人Id 
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`  //更新时间 
}


// TableName goldGoodsFile表 GoldGoodsFile自定义表名 gold_goods_file
func (GoldGoodsFile) TableName() string {
  return "gold_goods_file"
}

