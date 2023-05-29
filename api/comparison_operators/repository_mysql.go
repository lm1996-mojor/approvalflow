package comparison_operators

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

// SelectComparisonOperatorsListByIds 获取指定的分组信息列表
func (m MysqlRepository) SelectComparisonOperatorsListByIds(db *gorm.DB, ids []int64) (groupList []model.ComparisonOperators) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(ComparisonOperators.TableName()).Select(ComparisonOperators.GetAllColumn()).Scan(&groupList)
	return groupList
}

// SelectPageComparisonOperatorsByCondition 根据条件获取比较符列表
func (m MysqlRepository) SelectPageComparisonOperatorsByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ComparisonOperators) {
	db = db.Table(ComparisonOperators.TableName())
	db = db.Order("created_at asc")
	db = db.Where("classify_id = ?", mdl.ClassifyId).Select(ComparisonOperators.GetAllColumn()).Scan(&items)
	return
}

// SelectComparisonOperatorsById 获取单个比较符的信息
func (m MysqlRepository) SelectComparisonOperatorsById(db *gorm.DB, id int64) (processGroupsInfo model.ComparisonOperators) {
	db.Table(ComparisonOperators.TableName()).Where("id = ?", id).Select(ComparisonOperators.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertComparisonOperators 新增比较符
func (m MysqlRepository) InsertComparisonOperators(tx *gorm.DB, info model.ComparisonOperators) {
	if err := tx.Table(ComparisonOperators.TableName()).Create(&info).Error; err != nil {
		log.Error("新增比较符信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateComparisonOperatorsById 根据id修改比较符信息
func (m MysqlRepository) UpdateComparisonOperatorsById(tx *gorm.DB, info model.ComparisonOperators) {
	if err := tx.Table(ComparisonOperators.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改比较符信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteComparisonOperatorsById 删除比较符
func (m MysqlRepository) DeleteComparisonOperatorsById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ComparisonOperators.TableName()).Delete(&ComparisonOperators, ids).Error; err != nil {
		log.Error("删除比较符信息出错")
		panic("删除出错，数据已回滚")
	}
}
