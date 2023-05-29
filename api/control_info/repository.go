package control_info

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectCtlList(db *gorm.DB, mdl ListQueryMdl) (items []CtlDetail) //根据条件查询控件列表
	SelectSingleCtlInfo(db *gorm.DB, id int64) (ctlInfo CtlDetail)   //查询单个控件信息
	InsertCtl(tx *gorm.DB, info []model.ControlInfo)                 //批量新增控件
	UpdateCtl(tx *gorm.DB, info model.ControlInfo)                   //修改控件
}
