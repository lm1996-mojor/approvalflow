package repoImpl

import (
	"five.com/lk_flow/api/common"
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"gorm.io/gorm"
)

type CtlValueRepositoryImpl struct {
}

func CtlValueNewMysqlRepository() repository.CtlValueRepository {
	return CtlValueRepositoryImpl{}
}

func (m CtlValueRepositoryImpl) BatchInsertCtlValueInfos(tx *gorm.DB, ctlDetailList []api_model.CtlDetail) {
	ctlValueList := make([]model.CtlValueInfo, len(ctlDetailList))
	for i := 0; i < len(ctlValueList); i++ {
		ctlValueList[i].CtlValue = ctlDetailList[i].CtlValue
		ctlValueList[i].CtlId = ctlDetailList[i].CtlId
		ctlValueList[i].ApprovalCode = ctlDetailList[i].ApprovalCode
	}
	if err := tx.Table(common.CtlValueInfo.TableName()).CreateInBatches(&ctlValueList, 200).Error; err != nil {
		log.Error("新增流程表单控件值信息出错")
		panic("服务器错误")
	}
}

// SelectCtlValueListByApprovalCode 根据审批编号获取指定的审批表单
func (m CtlValueRepositoryImpl) SelectCtlValueListByApprovalCode(db *gorm.DB, approvalCode string) (ctlDetailList []api_model.CtlDetail) {
	db = db.Table(common.CtlValueInfo.TableName()+" ctv").Joins("LEFT JOIN "+common.ControlInfo.TableName()+" ctl on ctl.id=ctv.ctl_id").
		Where("ctv.approval_code = ?", approvalCode).Order("ctl.order_no asc")
	ctlSqlSelect := "ctl.parent_id,ctl.tab_id,ctl.owner_type,ctl.cn_name,ctl.en_name,ctl.ctl_code,ctl.enable," +
		"ctl.required,ctl.field_name,ctl.component_type,ctl.value_type,ctl.props,ctl.order_no,ctl.is_default"
	db.Select(common.CtlValueInfo.GetAllColumWithAlias("ctv") + " ," + ctlSqlSelect).Scan(&ctlDetailList)
	return ctlDetailList
}
