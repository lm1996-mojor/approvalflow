package repoImpl

import (
	"five.com/lk_flow/api/common"
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

func (m CtlValueRepositoryImpl) BatchInsertCtlValueInfos(tx *gorm.DB, ctlValueInfos []model.CtlValueInfo) {
	if err := tx.Table(common.CtlValueInfo.TableName()).CreateInBatches(&ctlValueInfos, 20).Error; err != nil {
		log.Error("新增流程表单控件值信息出错")
		panic("服务器错误")
	}
}
