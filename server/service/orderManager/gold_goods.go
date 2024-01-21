package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
    orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
)

type GoldGoodsService struct {
}

// CreateGoldGoods 创建goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService) CreateGoldGoods(goldGoods *orderManager.GoldGoods) (err error) {
	err = global.GVA_DB.Create(goldGoods).Error
	return err
}

// DeleteGoldGoods 删除goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService)DeleteGoldGoods(id string) (err error) {
	err = global.GVA_DB.Delete(&orderManager.GoldGoods{},"id = ?",id).Error
	return err
}

// DeleteGoldGoodsByIds 批量删除goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService)DeleteGoldGoodsByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]orderManager.GoldGoods{},"id in ?",ids).Error
	return err
}

// UpdateGoldGoods 更新goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService)UpdateGoldGoods(goldGoods orderManager.GoldGoods) (err error) {
	err = global.GVA_DB.Save(&goldGoods).Error
	return err
}

// GetGoldGoods 根据id获取goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService)GetGoldGoods(id string) (goldGoods orderManager.GoldGoods, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&goldGoods).Error
	return
}

// GetGoldGoodsInfoList 分页获取goldGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldGoodsService *GoldGoodsService)GetGoldGoodsInfoList(info orderManagerReq.GoldGoodsSearch) (list []orderManager.GoldGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&orderManager.GoldGoods{})
    var goldGoodss []orderManager.GoldGoods
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&goldGoodss).Error
	return  goldGoodss, total, err
}
