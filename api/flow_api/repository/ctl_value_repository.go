package repository

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// CtlValueRepository 操作数据库接口
type CtlValueRepository interface {
	BatchInsertCtlValueInfos(tx *gorm.DB, ctlValueInfos []model.CtlValueInfo) // 批量新增控件值信息
}
