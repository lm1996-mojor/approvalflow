package comparison_value

import (
	. "five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectComparisonValueListByIds 获取指定的分组信息列表
func (m MysqlRepository) SelectComparisonValueListByIds(db *gorm.DB, ids []int64) (groupList []model.ComparisonValue) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(ComparisonValue.TableName()).Select(ComparisonValue.GetAllColumn()).Scan(&groupList)
	return groupList
}

// SelectPageComparisonValueByCondition 根据条件获取比较值列表
func (m MysqlRepository) SelectPageComparisonValueByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ComparisonValue) {
	db = db.Table(ComparisonValue.TableName())
	db = db.Where("cond_detail_info_id = ?", mdl.CondDetailInfoId)
	db.Select(ComparisonValue.GetAllColumn()).Scan(&items)
	return
}

// SelectComparisonValueById 获取单个比较值的信息
func (m MysqlRepository) SelectComparisonValueById(db *gorm.DB, id int64) (processGroupsInfo model.ComparisonValue) {
	db.Table(ComparisonValue.TableName()).Where("id = ?", id).Select(ComparisonValue.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertComparisonValue 新增比较值
func (m MysqlRepository) InsertComparisonValue(tx *gorm.DB, info model.ComparisonValue) {
	if err := tx.Table(ComparisonValue.TableName()).Create(&info).Error; err != nil {
		log.Error("新增比较值信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateComparisonValueById 根据id修改比较值信息
func (m MysqlRepository) UpdateComparisonValueById(tx *gorm.DB, info model.ComparisonValue) {
	if err := tx.Table(ComparisonValue.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改比较值信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteComparisonValueById 删除比较值
func (m MysqlRepository) DeleteComparisonValueById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ComparisonValue.TableName()).Delete(&ComparisonValue, ids).Error; err != nil {
		log.Error("删除比较值信息出错")
		panic("删除出错，数据已回滚")
	}
}
