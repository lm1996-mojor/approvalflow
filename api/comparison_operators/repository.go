package comparison_operators

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectComparisonOperatorsListByIds(db *gorm.DB, ids []int64) []model.ComparisonOperators                    //获取指定的分组信息
	SelectPageComparisonOperatorsByCondition(db *gorm.DB, mdl QueryListMdl) (items []model.ComparisonOperators) //根据条件获取比较符列表
	SelectComparisonOperatorsById(db *gorm.DB, id int64) model.ComparisonOperators                              //获取单个比较符的信息
	InsertComparisonOperators(tx *gorm.DB, info model.ComparisonOperators)                                      //新增比较符
	UpdateComparisonOperatorsById(tx *gorm.DB, info model.ComparisonOperators)                                  //根据id修改比较符信息
	DeleteComparisonOperatorsById(tx *gorm.DB, ids []int64)                                                     //删除比较符
}
