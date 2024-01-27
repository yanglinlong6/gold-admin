package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      `gorm:"autoCreateTime"`       // 创建时间
	UpdatedAt time.Time      `gorm:"autoCreateTime"`       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
