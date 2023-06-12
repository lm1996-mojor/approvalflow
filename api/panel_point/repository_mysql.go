package panel_point

import (
	"five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/lk_flow/model/rel"
	"five.com/technical_center/core_library.git/log"
	lib "five.com/technical_center/core_library.git/utils/repo"
	"five.com/technical_center/core_library.git/utils/trans"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectPanelPointListByCondition 根据条件获取流程节点列表
func (m MysqlRepository) SelectPanelPointListByCondition(db *gorm.DB, condition map[string]interface{}) (panelPointList []PanelPointDetail) {
	var panelPoints []model.PanelPoint
	for k, v := range condition {
		db = db.Where(k, v)
	}
	db.Table(common.PanelPoint.TableName()).Select(common.PanelPoint.GetAllColumn()).Scan(&panelPoints)
	trans.DeepCopy(panelPoints, &panelPointList)
	for i := 0; i < len(panelPoints); i++ {
		selectPreviousStepConditionMap := make(map[string]interface{}, 0)
		selectPreviousStepConditionMap["point_id = ?"] = panelPoints[i].Id
		relMdlInfos := selectRelInfoListByCondition(lib.ObtainCustomDb(), selectPreviousStepConditionMap)
		trans.DeepCopy(relMdlInfos, &panelPointList[i].RelMdl)
	}
	return panelPointList
}

// SelectSpecifyNodesNextStepInfoByPointId 根据指定的节点id查询其下一个节点的信息
func (m MysqlRepository) SelectSpecifyNodesNextStepInfoByPointId(db *gorm.DB, pointId int64) (panelPointList []model.PanelPoint) {
	selectPreviousStepConditionMap := make(map[string]interface{}, 0)
	selectPreviousStepConditionMap["point_id = ?"] = pointId
	relInfos := selectRelInfoListByCondition(db, selectPreviousStepConditionMap)
	for _, relMdl := range relInfos {
		panelPointList = append(panelPointList, m.SelectSinglePanelPointInfoById(lib.ObtainCustomDb(), relMdl.NextStep))
	}
	return panelPointList
}

// 根据条件查询节点关系描述列表
func selectRelInfoListByCondition(db *gorm.DB, conditionMap map[string]interface{}) (relMdlInfos []rel.PointRelDesc) {
	db = db.Table(common.PointRelDesc.TableName())
	db = conditionBuilder(db, conditionMap)
	if err := db.Select(common.PointRelDesc.GetAllColumn()).Scan(&relMdlInfos).Error; err != nil {
		log.Error("查询单个节点关系描述错误,请检查查询条件是否正确")
		panic("查询错误，联系管理员")
	}
	return
}

// SelectSinglePanelPointInfoById 获取单个流程节点信息
func (m MysqlRepository) SelectSinglePanelPointInfoById(db *gorm.DB, id int64) (info model.PanelPoint) {
	db.Table(common.PanelPoint.TableName()).Where("id = ?", id).Select(common.PanelPoint.GetAllColumn()).Scan(&info)
	return info
}

// SelectSinglePanelPointInfoByCondition 根据条件获取单个节点及其关系描述信息
func (m MysqlRepository) SelectSinglePanelPointInfoByCondition(db *gorm.DB, condition map[string]interface{}) (panelPointList PanelPointDetail) {
	var panelPoint model.PanelPoint
	for k, v := range condition {
		db = db.Where(k, v)
	}
	db.Table(common.PanelPoint.TableName()).Select(common.PanelPoint.GetAllColumn()).Scan(&panelPoint)
	trans.DeepCopy(panelPoint, &panelPointList)
	//查询描述关系
	selectPreviousStepConditionMap := make(map[string]interface{}, 0)
	selectPreviousStepConditionMap["point_id = ?"] = panelPoint.Id
	relMdlInfos := selectRelInfoListByCondition(lib.ObtainCustomDb(), selectPreviousStepConditionMap)
	trans.DeepCopy(relMdlInfos, &panelPointList.RelMdl)
	return panelPointList
}

// InsertPanelPoint 新增流程节点
func (m MysqlRepository) InsertPanelPoint(tx *gorm.DB, info PanelPointDetail) int64 {
	var panelPointInfo model.PanelPoint
	trans.CopyFields(info, &panelPointInfo)
	//判断是否为条件节点

	if err := tx.Table(common.PanelPoint.TableName()).Create(&panelPointInfo).Error; err != nil {
		log.Error("新增流程节点信息出错")
		panic("服务器错误")
	}

	//节点类型为条件节点时不会记录该节点的关系，关系记录会与条件主体产生关联
	createRelLogic(tx, info.RelMdl, panelPointInfo.Id, panelPointInfo.PointType, false)
	if panelPointInfo.PointType == 4 {
		//先判断该流程中是否已经存在条件节点,如果已有条件节点则不需要添加默认条件节点，如果没有则添加默认条件节点
		var conditionPointTotal int64
		lib.ObtainCustomDb().Table(common.PanelPoint.TableName()).Where("process_id = ?", info.ProcessId).Where("point_type = ?", 4).Count(&conditionPointTotal)
		if conditionPointTotal <= 1 {
			var pointInfo model.PanelPoint
			trans.DeepCopy(info, &pointInfo)
			pointInfo.PointName = "默认条件"
			pointInfo.Priority = 100
			pointInfo.ConditionType = 2
			pointInfo.PointType = 4
			pointInfo.ProcessId = panelPointInfo.ProcessId
			if err := tx.Table(common.PanelPoint.TableName()).Create(&pointInfo).Error; err != nil {
				log.Error("新增流程节点信息出错")
				panic("服务器错误")
			}
			// 创建关系描述
			createRelLogic(tx, info.RelMdl, pointInfo.Id, pointInfo.PointType, true)
		} else {
			// 默认条件的优先级永远最小，所以需要判断当前新增的节点的优先级是否大于目前流程中优先级最低的条件节点
			// 同层级中,优先级最小的条件节点
			var minPriorityPointInfo model.PanelPoint
			// 子查询
			db := lib.ObtainCustomDb().Table(common.PanelPoint.TableName()).Where("process_id = ?", info.ProcessId).Where("point_type = ?", 4)
			db = db.Where("priority = (?)", lib.ObtainCustomDb().Table(common.PanelPoint.TableName()+" p").Joins("left join point_rel_desc prd on prd.point_id=p.id").
				Where("prd.previous_step = ?", info.RelMdl[0].PreviousStep).
				Where("p.process_id = ?", panelPointInfo.ProcessId).
				Where("p.point_type = ?", 4).
				Where("p.condition_type = ?", 2).
				Select("MAX(p.priority)"))
			db.Scan(&minPriorityPointInfo)
			//判断当前节点的优先级是否大于等于最小的优先级
			if panelPointInfo.Priority >= minPriorityPointInfo.Priority {
				minPriorityPointInfo.Priority = panelPointInfo.Priority + 1
				m.UpdatePanelPoint(tx, minPriorityPointInfo)
			}
		}
	}

	return info.Id
}

// 创建节点关系描述信息
func createRelInfo(tx *gorm.DB, relMdl rel.PointRelDesc) {
	if err := tx.Table(common.PointRelDesc.TableName()).Create(&relMdl).Error; err != nil {
		log.Error("新增流程节点关系描述信息出错")
		panic("服务器错误")
	}

}

// 创建节点关系信息
func createRelLogic(tx *gorm.DB, relMdls []rel.PointRelDesc, pointId int64, pointType int8, pointIsFirstAdd bool) {
	if len(relMdls) > 1 {
		if relMdls[0].NextStep == relMdls[1].NextStep {
			var nextStepRelMdlList []rel.PointRelDesc
			lib.ObtainCustomDb().Table(common.PointRelDesc.TableName()).Where("point_id = ?", relMdls[0].NextStep).Scan(&nextStepRelMdlList)
			for i := 1; i < len(nextStepRelMdlList); i++ {
				conditionParams := make(map[string]interface{}, 0)
				conditionParams["id = ?"] = nextStepRelMdlList[i].Id
				deletePanelPointRelDescByConditionOnIds(tx, conditionParams)
			}
		}
	}
	for _, desc := range relMdls {
		//var relMdl rel.PointRelDesc
		//trans.DeepCopy(desc, &relMdl)
		desc.PointId = pointId
		//创建关系描述
		createRelInfo(tx, desc)
		//判断是否为第一次新增节点(即插入节点)
		if pointIsFirstAdd {
			relMaintenanceLogic(tx, desc, pointId, pointType)
		} else {
			// 如果该节点第一次新增且该节点为条件节点，则为新增一个条件。则会为该节点的上下级节点进行关系描述信息新增加一条
			if pointType == 4 {
				insertUpAndDownPointRelLogic(tx, desc, pointId, pointType)
			} else {
				relMaintenanceLogic(tx, desc, pointId, pointType)
			}
		}
	}
}

// 关联上下节点的关系-逻辑(主要为了封装，减少代码量)
func relMaintenanceLogic(tx *gorm.DB, desc rel.PointRelDesc, pointId int64, pointType int8) {
	//连接下一个节点
	if desc.NextStep != 0 {
		conditionParams := make(map[string]interface{}, 0)
		conditionParams["point_id = ?"] = desc.NextStep
		if desc.PreviousStep != 0 {
			conditionParams["previous_step = ?"] = desc.PreviousStep
			conditionParams["previous_step_type = ?"] = desc.PreviousStepType
		}
		updateRelMdlByCondition(tx,
			conditionParams,
			map[string]interface{}{"previous_step": pointId, "previous_step_type": pointType})

	}
	//连接上一个节点
	if desc.PreviousStep != 0 {
		conditionParams := make(map[string]interface{}, 0)
		conditionParams["point_id = ?"] = desc.PreviousStep
		if desc.NextStep != 0 {
			conditionParams["next_step = ?"] = desc.NextStep
			conditionParams["next_step_type = ?"] = desc.NextStepType
		}
		updateRelMdlByCondition(tx,
			conditionParams,
			map[string]interface{}{"next_step": pointId, "next_step_type": pointType})
	}
}

// 新增上下节点中的关系描述-逻辑(主要为了封装，减少代码量)
func insertUpAndDownPointRelLogic(tx *gorm.DB, desc rel.PointRelDesc, pointId int64, pointType int8) {
	log.Info("测试")
	// 上一个节点新增关系
	var previousStepRelMdl rel.PointRelDesc
	previousStepRelMdls := selectRelInfoListByCondition(lib.ObtainCustomDb(), map[string]interface{}{"point_id": desc.PreviousStep})
	// 使用复制会出现主键无法自增情况
	previousStepRelMdl.PreviousStep = previousStepRelMdls[0].PreviousStep
	previousStepRelMdl.PreviousStepType = previousStepRelMdls[0].PreviousStepType
	previousStepRelMdl.PointId = previousStepRelMdls[0].PointId
	previousStepRelMdl.NextStep = pointId
	previousStepRelMdl.NextStepType = pointType
	//创建关系描述
	createRelInfo(tx, previousStepRelMdl)
	// 下一个节点新增关系
	var nextStepRelMdl rel.PointRelDesc
	nextStepRelMdls := selectRelInfoListByCondition(lib.ObtainCustomDb(), map[string]interface{}{"point_id": desc.NextStep})
	// 使用复制会出现主键无法自增情况
	nextStepRelMdl.PreviousStep = pointId
	nextStepRelMdl.PreviousStepType = pointType
	nextStepRelMdl.PointId = nextStepRelMdls[0].PointId
	nextStepRelMdl.NextStep = nextStepRelMdls[0].NextStep
	nextStepRelMdl.NextStepType = nextStepRelMdls[0].NextStepType
	//创建关系描述
	createRelInfo(tx, nextStepRelMdl)
}

// UpdatePanelPoint 根据id修改流程节点信息
func (m MysqlRepository) UpdatePanelPoint(tx *gorm.DB, info model.PanelPoint) {
	if err := tx.Table(common.PanelPoint.TableName()).Updates(&info).Error; err != nil {
		log.Error("修改流程节点信息出错")
		panic("修改出错，数据已回滚")
	}
}

// 根据条件修改节点信息
func updatePanelPointByCondition(tx *gorm.DB, conditionMap map[string]interface{}, updateColumns map[string]interface{}) {
	tx = tx.Table(common.PanelPoint.TableName())
	tx = conditionBuilder(tx, conditionMap)
	if err := tx.Updates(&updateColumns).Error; err != nil {
		log.Error("修改流程节点信息出错")
		panic("修改出错，数据已回滚")
	}
}

// 根据条件更新节点关系描述信息
func updateRelMdlByCondition(tx *gorm.DB, conditionMap map[string]interface{}, updateColumns map[string]interface{}) {
	tx = tx.Table(common.PointRelDesc.TableName())
	tx = conditionBuilder(tx, conditionMap)
	if err := tx.Updates(&updateColumns).Error; err != nil {
		log.Error("修改节点关系信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeletePanelPoint 删除流程节点
func (m MysqlRepository) DeletePanelPoint(tx *gorm.DB, ids []int64) {
	//修改当前删除节点的上下节点的关系
	for _, id := range ids {
		// 当前节点的信息
		panelPointInfo := m.SelectSinglePanelPointInfoById(lib.ObtainCustomDb(), id)
		// 修改前端所需的字段值:删除当前节点前，先将当前节点的下一个节点,中上一个节点的值改为当前节点的上一个节点id
		/*
				数据示例：
			      原数据 ->
			 	 	A    id: 1     	only_web_use_previous_step: 0
					B    id: 2     	only_web_use_previous_step: 1
					C	 id: 3		only_web_use_previous_step: 2
			      删除数据 ->   删除节点 B
					(不动)A    id: 1     only_web_use_previous_step: 0
					(不动)C	  id: 3		only_web_use_previous_step: 1
		*/

		// 当前节点所包含的所有关系信息
		selectPreviousStepConditionMap := make(map[string]interface{}, 0)
		selectPreviousStepConditionMap["point_id = ?"] = panelPointInfo.Id
		relMdlInfos := selectRelInfoListByCondition(lib.ObtainCustomDb(), selectPreviousStepConditionMap)
		for _, info := range relMdlInfos {
			//修改上一节点
			updatePreviousStepConditionMap := make(map[string]interface{}, 0)
			updatePreviousStepConditionMap["next_step = ?"] = panelPointInfo.Id
			updatePreviousStepConditionMap["next_step_type = ?"] = panelPointInfo.PointType
			updateRelMdlByCondition(tx,
				updatePreviousStepConditionMap,
				map[string]interface{}{"next_step": info.NextStep, "next_step_type": info.NextStepType})
			//修改下一节点
			updateNextStepConditionMap := make(map[string]interface{}, 0)
			updateNextStepConditionMap["previous_step = ?"] = panelPointInfo.Id
			updateNextStepConditionMap["previous_step_type = ?"] = panelPointInfo.PointType
			updateRelMdlByCondition(tx,
				updatePreviousStepConditionMap,
				map[string]interface{}{"previous_step": info.PreviousStep, "previous_step_type": info.PreviousStepType})
		}
	}
	//删除节点关系描述
	deleteConditionMap := make(map[string]interface{}, 0)
	deleteConditionMap["point_id in (?)"] = ids
	deletePanelPointRelDescByConditionOnIds(tx, deleteConditionMap)
	//删除节点信息
	if err := tx.Table(common.PanelPoint.TableName()).Unscoped().Delete(&common.PanelPoint, ids).Error; err != nil {
		log.Error("删除流程节点信息出错")
		panic("删除出错，数据已回滚")
	}
}

// 根据条件删除节点关系描述信息
func deletePanelPointRelDescByConditionOnIds(tx *gorm.DB, conditionMap map[string]interface{}) {
	tx.Table(common.PointRelDesc.TableName()).Unscoped()
	tx = conditionBuilder(tx, conditionMap)
	if err := tx.Delete(&common.PointRelDesc).Error; err != nil {
		log.Error("删除流程节点关系描述信息出错")
		panic("删除出错，数据已回滚")
	}
}

// 查询条件构造器
func conditionBuilder(obj *gorm.DB, conditionMap map[string]interface{}) *gorm.DB {
	for k, v := range conditionMap {
		obj = obj.Where(k, v)
	}
	return obj
}
