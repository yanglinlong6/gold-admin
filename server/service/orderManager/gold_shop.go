package orderManager

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orderManager"
	orderManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/request"
	orderManagerResp "github.com/flipped-aurora/gin-vue-admin/server/model/orderManager/response"
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type GoldShopService struct {
}

// CreateGoldShop 创建goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) CreateGoldShop(goldShop *orderManager.GoldShop, userId int) (err error) {
	goldShop.Status = new(int)
	currentTime := time.Now()
	goldShop.CreateId = &userId
	goldShop.CreateTime = &currentTime
	err = global.GVA_DB.Create(goldShop).Error
	return err
}

// DeleteGoldShop 删除goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) DeleteGoldShop(id string) (err error) {
	err = global.GVA_DB.Delete(&orderManager.GoldShop{}, "id = ?", id).Error
	return err
}

// DeleteGoldShopByIds 批量删除goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) DeleteGoldShopByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]orderManager.GoldShop{}, "id in ?", ids).Error
	return err
}

// UpdateGoldShop 更新goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) UpdateGoldShop(goldShop orderManager.GoldShop, userId int) (err error) {
	// 通过ID获取记录
	goldShopResp := orderManager.GoldShop{}
	if err = global.GVA_DB.Where("id = ?", goldShop.ID).First(&goldShopResp).Error; err != nil {
		return err
	}
	goldShopResp.GoodsId = goldShop.GoodsId
	goldShopResp.BuyId = goldShop.BuyId
	goldShopResp.Status = goldShop.Status
	currentTime := time.Now()
	goldShopResp.UpdateId = &userId
	goldShopResp.UpdateTime = &currentTime
	err = global.GVA_DB.Save(&goldShopResp).Error
	return err
}

// GetGoldShop 根据id获取goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) GetGoldShop(id string) (GoldShopResp orderManagerResp.GoldShopResp, err error) {
	goldShop := orderManager.GoldShop{}
	err = global.GVA_DB.Where("id = ?", id).First(&goldShop).Error
	sysUserList := []systemModel.SysUser{}
	global.GVA_DB.Find(&sysUserList)
	sysUserMap := make(map[uint]systemModel.SysUser)
	for _, element := range sysUserList {
		sysUserMap[element.ID] = element
	}
	if goldShop.BuyId == nil {
		goldShop.BuyId = new(int)
	}
	if goldShop.CreateId == nil {
		goldShop.CreateId = new(int)
	}
	if goldShop.UpdateId == nil {
		goldShop.UpdateId = new(int)
	}

	// 根据商品ID获取商品名称
	var goods orderManager.GoldGoods
	global.GVA_DB.Where("id = ?", goldShop.GoodsId).First(&goods)

	goldShopResp := orderManagerResp.GoldShopResp{
		ID:         goldShop.ID,
		GoodsId:    goldShop.GoodsId,
		GoodsName:  goods.GoodsName,
		BuyId:      goldShop.BuyId,
		BuyName:    sysUserMap[uint(*goldShop.BuyId)].Username,
		Status:     goldShop.Status,
		CreateName: sysUserMap[uint(*goldShop.CreateId)].Username,
		CreateTime: goldShop.CreateTime,
		UpdateName: sysUserMap[uint(*goldShop.CreateId)].Username,
		UpdateTime: goldShop.UpdateTime,
	}
	return goldShopResp, err
}

// GetGoldShopInfoList 分页获取goldShop表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldShopService *GoldShopService) GetGoldShopInfoList(info orderManagerReq.GoldShopSearch) (list []orderManagerResp.GoldShopResp, total int64, err error) {
	sysUserList := []systemModel.SysUser{}
	global.GVA_DB.Find(&sysUserList)
	sysUserMap := make(map[uint]systemModel.SysUser)
	for _, element := range sysUserList {
		sysUserMap[element.ID] = element
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&orderManager.GoldShop{})
	var goldShops []orderManager.GoldShop
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&goldShops).Error
	goldShopRespList := []orderManagerResp.GoldShopResp{}
	for _, element := range goldShops {
		if element.BuyId == nil {
			element.BuyId = new(int)
		}
		if element.CreateId == nil {
			element.CreateId = new(int)
		}
		if element.UpdateId == nil {
			element.UpdateId = new(int)
		}

		// 根据商品ID获取商品名称
		var goods orderManager.GoldGoods
		global.GVA_DB.Where("id = ?", element.GoodsId).First(&goods)

		goldShopResp := orderManagerResp.GoldShopResp{
			ID:         element.ID,
			GoodsId:    element.GoodsId,
			GoodsName:  goods.GoodsName,
			BuyId:      element.BuyId,
			BuyName:    sysUserMap[uint(*element.BuyId)].Username,
			Status:     element.Status,
			CreateName: sysUserMap[uint(*element.CreateId)].Username,
			CreateTime: element.CreateTime,
			UpdateName: sysUserMap[uint(*element.CreateId)].Username,
			UpdateTime: element.UpdateTime,
		}
		goldShopRespList = append(goldShopRespList, goldShopResp)
	}
	return goldShopRespList, total, err
}
