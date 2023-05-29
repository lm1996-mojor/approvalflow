package comparison_value

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectComparisonValueListByIds(db *gorm.DB, ids []int64) []model.ComparisonValue                    //获取指定的分组信息
	SelectPageComparisonValueByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ComparisonValue) //根据条件获取比较值列表
	SelectComparisonValueById(db *gorm.DB, id int64) model.ComparisonValue                              //获取单个比较值的信息
	InsertComparisonValue(tx *gorm.DB, info model.ComparisonValue)                                      //新增比较值
	UpdateComparisonValueById(tx *gorm.DB, info model.ComparisonValue)                                  //根据id修改比较值信息
	DeleteComparisonValueById(tx *gorm.DB, ids []int64)                                                 //删除比较值
}
