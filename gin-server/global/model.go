package global

import (
	"time"
)

type MODEL struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT;comment:主键ID"` // 主键ID
	CreatedAt time.Time      `gorm:"comment:创建时间"`// 创建时间
	UpdatedAt time.Time      `gorm:"comment:更新时间"`// 更新时间
	DeletedAt time.Time 	`gorm:"comment:删除时间"`// 删除时间
}
