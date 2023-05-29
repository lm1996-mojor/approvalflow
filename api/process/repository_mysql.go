package process

import (
	. "five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/lk_flow/model/rel"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectProcessListByCondition 根据条件获取流程列表
func (m MysqlRepository) SelectProcessListByCondition(db *gorm.DB, list []ListProcessDetail, mdl QueryListMdl) []ListProcessDetail {
	for i := 0; i < len(list); i++ {
		var processList []model.Process
		if mdl.Search != "" || len(mdl.Search) != 0 {
			db = db.Where("flow_name like ?", "%"+mdl.Search+"%")
		}
		db = db.Where("group_id = ?", list[i].Id)
		db.Table(Process.TableName()).Select(Process.GetAllColumn()).Scan(&processList)
		trans.DeepCopy(processList, list[i].ProcessList)
	}
	return list
}

// SelectProcessInfo 获取单个流程信息
func (m MysqlRepository) SelectProcessInfo(db *gorm.DB, id int64) (processInfo DetailProcess) {
	db.Table(Process.TableName()+" p").
		Joins("Left Join process_groups pg on pg.id = p.group_id").
		Where("p.id = ?", id).
		Select(Process.GetAllColumWithAlias("p") + ", pg.group_name").
		Scan(&processInfo)
	return processInfo
}

// InsertProcess 新增流程
func (m MysqlRepository) InsertProcess(tx *gorm.DB, info model.Process) int64 {
	//创建流程
	if err := tx.Table(Process.TableName()).Create(&info).Error; err != nil {
		log.Error("新增流程信息出错")
		panic("服务器错误")
	}

	// 创建发起人节点和结束节点
	var panelPointInfoStart model.PanelPoint
	panelPointInfoStart.PointType = 5
	panelPointInfoStart.PointName = "发起人节点"
	panelPointInfoStart.ProcessId = info.Id
	if err := tx.Table(PanelPoint.TableName()).Create(&panelPointInfoStart).Error; err != nil {
		log.Error("新增流程发起人节点信息出错")
		panic("服务器错误")
	}

	var panelPointInfoEnd model.PanelPoint
	panelPointInfoEnd.PointType = 6
	panelPointInfoEnd.PointName = "结束节点"
	panelPointInfoEnd.OnlyWebUsePreviousStep = panelPointInfoStart.Id
	panelPointInfoEnd.ProcessId = info.Id
	if err := tx.Table(PanelPoint.TableName()).Create(&panelPointInfoEnd).Error; err != nil {
		log.Error("新增流程结束节点信息出错")
		panic("服务器错误")
	}

	// 创建两个节点并连接
	var startPointRelMdl rel.PointRelDesc
	startPointRelMdl.PointId = panelPointInfoStart.Id
	startPointRelMdl.NextStep = panelPointInfoEnd.Id
	startPointRelMdl.NextStepType = panelPointInfoEnd.PointType
	if err := tx.Table(PointRelDesc.TableName()).Create(&startPointRelMdl).Error; err != nil {
		log.Error("新增流程节点关系描述信息出错")
		panic("服务器错误")
	}

	var endPointRelMdl rel.PointRelDesc
	endPointRelMdl.PointId = panelPointInfoEnd.Id
	endPointRelMdl.PreviousStep = panelPointInfoStart.Id
	endPointRelMdl.PreviousStepType = panelPointInfoStart.PointType
	if err := tx.Table(PointRelDesc.TableName()).Create(&endPointRelMdl).Error; err != nil {
		log.Error("新增流程节点关系描述信息出错")
		panic("服务器错误")
	}
	return info.Id
}

// UpdateProcess 根据id修改流程信息
func (m MysqlRepository) UpdateProcess(tx *gorm.DB, info model.Process) {
	if err := tx.Table(Process.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改流程信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteProcess 删除流程
func (m MysqlRepository) DeleteProcess(tx *gorm.DB, ids []int64) {
	if err := tx.Table(Process.TableName()).Delete(&Process, ids).Error; err != nil {
		log.Error("删除流程信息出错")
		panic("删除出错，数据已回滚")
	}
}

// SelectProcessListByIds 获取指定流程主体信息
func (m MysqlRepository) SelectProcessListByIds(db *gorm.DB, ids []int64) (processList []model.Process) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(Process.TableName()).Select(Process.GetAllColumn()).Scan(&processList)
	return processList
}

// UpdateProcessSpecifyColumns 更新单个流程指定列
func (m MysqlRepository) UpdateProcessSpecifyColumns(tx *gorm.DB, id int64, columns map[string]interface{}) {
	if err := tx.Table(Process.TableName()).Where("id = ?", id).Updates(columns).Error; err != nil {
		log.Error("发布流程出错")
		panic("修改出错，数据已回滚")
	}
}
