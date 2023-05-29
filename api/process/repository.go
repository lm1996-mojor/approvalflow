package process

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectProcessListByCondition(db *gorm.DB, list []ListProcessDetail, mdl QueryListMdl) []ListProcessDetail //根据条件获取流程列表
	SelectProcessInfo(db *gorm.DB, id int64) DetailProcess                                                    //获取单个流程信息
	InsertProcess(tx *gorm.DB, info model.Process) int64                                                      //新增流程
	UpdateProcess(tx *gorm.DB, info model.Process)                                                            //根据id修改流程信息
	DeleteProcess(tx *gorm.DB, ids []int64)                                                                   //删除流程
	SelectProcessListByIds(db *gorm.DB, ids []int64) (processList []model.Process)                            //获取指定流程主体信息
	UpdateProcessSpecifyColumns(tx *gorm.DB, id int64, columns map[string]interface{})                        //更新单个流程指定列
}
