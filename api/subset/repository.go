package subset

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectSubList(db *gorm.DB, mdl ListQueryMdl) []SubDetail                          //根据条件查询子集列表（分页）
	SelectSubInfo(db *gorm.DB, id int64) SubDetail                                    //查询单个子集信息
	InsertSub(tx *gorm.DB, info model.Subset)                                         //新增子集
	UpdateSub(tx *gorm.DB, info SubDetail)                                            //修改子集
	DeleteSub(tx *gorm.DB, ids []int64)                                               //删除子集
	UpdateSubSpecifyColumnsById(db *gorm.DB, id int64, params map[string]interface{}) //根据id修改子集信息(指定列)
	SelectSubsetInfosByParentId(db *gorm.DB, id int64) []model.Subset                 //根据子集父级id获取多个子集信息
}
