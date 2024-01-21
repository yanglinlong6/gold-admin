package orderManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
    orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
)

type GoldShopService struct {
}

// CreateGoldShop 创建goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) CreateGoldShop(goldShop *orderManager.GoldShop) (err error) {
	err = global.GVA_DB.Create(goldShop).Error
	return err
}

// DeleteGoldShop 删除goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService)DeleteGoldShop(id string) (err error) {
	err = global.GVA_DB.Delete(&orderManager.GoldShop{},"id = ?",id).Error
	return err
}

// DeleteGoldShopByIds 批量删除goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService)DeleteGoldShopByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]orderManager.GoldShop{},"id in ?",ids).Error
	return err
}

// UpdateGoldShop 更新goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService)UpdateGoldShop(goldShop orderManager.GoldShop) (err error) {
	err = global.GVA_DB.Save(&goldShop).Error
	return err
}

// GetGoldShop 根据id获取goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService)GetGoldShop(id string) (goldShop orderManager.GoldShop, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&goldShop).Error
	return
}

// GetGoldShopInfoList 分页获取goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService)GetGoldShopInfoList(info orderManagerReq.GoldShopSearch) (list []orderManager.GoldShop, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&orderManager.GoldShop{})
    var goldShops []orderManager.GoldShop
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
	
	err = db.Find(&goldShops).Error
	return  goldShops, total, err
}
