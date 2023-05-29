package subset

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

// SelectSubList 根据条件查询子集列表（分页）
func (m MysqlRepository) SelectSubList(db *gorm.DB, mdl ListQueryMdl) (items []SubDetail) {
	db.Table(Subset.TableName()).Where("process_id = ?").Scan(&items)
	return items
}

// SelectSubInfo 查询单个子集信息
func (m MysqlRepository) SelectSubInfo(db *gorm.DB, id int64) (subInfo SubDetail) {
	db.Table(Subset.TableName()+" ss").Where("ss.id = ?", id).Select(Subset.GetAllColumWithAlias("ss")).Scan(&subInfo)
	return subInfo
}

// InsertSub 新增子集
func (m MysqlRepository) InsertSub(tx *gorm.DB, info model.Subset) {
	if err := tx.Table(Subset.TableName()).Create(&info).Error; err != nil {
		log.Error("新增控件子集出错，原因>>" + err.Error())
		panic("新增出错，数据已回滚")
	}
}

// UpdateSub 修改子集
func (m MysqlRepository) UpdateSub(tx *gorm.DB, info SubDetail) {
	if err := tx.Table(Subset.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改控件子集出错，原因>>" + err.Error())
		panic("修改出错，数据已回滚")
	}
}

// DeleteSub 删除子集
func (m MysqlRepository) DeleteSub(tx *gorm.DB, ids []int64) {
	if err := tx.Unscoped().Table(Subset.TableName()).Delete(&Subset, ids).Error; err != nil {
		log.Error("删除控件子集出错，原因>>" + err.Error())
		panic("删除出错，数据已回滚")
	}
}

// UpdateSubSpecifyColumnsById 根据id修改子集信息(指定列)
func (m MysqlRepository) UpdateSubSpecifyColumnsById(tx *gorm.DB, id int64, params map[string]interface{}) {
	if err := tx.Model(&Subset).Updates(params).Where("id = ?", id).Error; err != nil {
		log.Error("修改控件子集指定列出错，原因>>" + err.Error())
		panic("修改出错，数据已回滚")
	}
}

// SelectSubsetInfosByParentId 根据子集父级id获取多个子集信息
func (m MysqlRepository) SelectSubsetInfosByParentId(db *gorm.DB, id int64) (subsetInfos []model.Subset) {
	db.Table(Subset.TableName()).Where("parent_id = ?", id).Select(Subset.GetAllColumn()).Scan(&subsetInfos)
	return
}
