package comparison_classify

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectComparisonClassifyListByIds(db *gorm.DB, ids []int64) []model.ComparisonClassify                     //获取指定的分组信息
	SelectPageComparisonClassifyByCondition(db *gorm.DB, mdl QueryPageMdl) ([]model.ComparisonClassify, int64) //根据条件获取比较符分类列表
	SelectComparisonClassifyById(db *gorm.DB, id int64) model.ComparisonClassify                               //获取单个比较符分类的信息
	InsertComparisonClassify(tx *gorm.DB, info model.ComparisonClassify)                                       //新增比较符分类
	UpdateComparisonClassifyById(tx *gorm.DB, info model.ComparisonClassify)                                   //根据id修改比较符分类信息
	DeleteComparisonClassifyById(tx *gorm.DB, ids []int64)                                                     //删除比较符分类
	SelectAllComparisonClassifyList(db *gorm.DB) []model.ComparisonClassify                                    //获取所有比较符分类
}
