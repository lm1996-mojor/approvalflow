package comparison_classify

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

// SelectComparisonClassifyListByIds 获取指定的分组信息列表
func (m MysqlRepository) SelectComparisonClassifyListByIds(db *gorm.DB, ids []int64) (groupList []model.ComparisonClassify) {
	if len(ids) > 0 {
		db = db.Where("id in(?)", ids)
	}
	db.Table(ComparisonClassify.TableName()).Select(ComparisonClassify.GetAllColumn()).Scan(&groupList)
	return groupList
}

// SelectPageComparisonClassifyByCondition 根据条件获取比较符分类列表
func (m MysqlRepository) SelectPageComparisonClassifyByCondition(db *gorm.DB, mdl QueryPageMdl) (items []model.ComparisonClassify, total int64) {
	db = db.Table(ComparisonClassify.TableName())
	if mdl.Search != "" || len(mdl.Search) > 0 {
		db = db.Where("classify_name like ?", "%"+mdl.Search+"%")
	}
	db = db.Count(&total)
	db = db.Order("created_at asc")
	mdl.PageNumber = mdl.PageNumber * mdl.PageSize
	if mdl.PageSize <= 0 {
		mdl.PageSize = 10
	}
	db.Select(ComparisonClassify.GetAllColumn())
	db.Limit(mdl.PageSize).Offset(mdl.PageNumber).Scan(&items)
	return
}

// SelectComparisonClassifyById 获取单个比较符分类的信息
func (m MysqlRepository) SelectComparisonClassifyById(db *gorm.DB, id int64) (processGroupsInfo model.ComparisonClassify) {
	db.Table(ComparisonClassify.TableName()).Where("id = ?", id).Select(ComparisonClassify.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertComparisonClassify 新增比较符分类
func (m MysqlRepository) InsertComparisonClassify(tx *gorm.DB, info model.ComparisonClassify) {
	if err := tx.Table(ComparisonClassify.TableName()).Create(&info).Error; err != nil {
		log.Error("新增比较符分类信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateComparisonClassifyById 根据id修改比较符分类信息
func (m MysqlRepository) UpdateComparisonClassifyById(tx *gorm.DB, info model.ComparisonClassify) {
	if err := tx.Table(ComparisonClassify.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改比较符分类信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteComparisonClassifyById 删除比较符分类
func (m MysqlRepository) DeleteComparisonClassifyById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ComparisonClassify.TableName()).Delete(&ComparisonClassify, ids).Error; err != nil {
		log.Error("删除比较符分类信息出错")
		panic("删除出错，数据已回滚")
	}
}

// SelectAllComparisonClassifyList 获取所有比较符分类
func (m MysqlRepository) SelectAllComparisonClassifyList(db *gorm.DB) (comparisonClassifyList []model.ComparisonClassify) {
	db.Where(ComparisonClassify.TableName()).Select(ComparisonClassify.GetAllColumn()).Scan(&comparisonClassifyList)
	return
}
