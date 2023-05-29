package comparison_classify

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainComparisonClassifyPage(mdl QueryPageMdl) rest.Result          //根据条件获取比较符分类列表
	ObtainComparisonClassifyInfo(id int64) rest.Result                  //获取单个比较符分类的信息
	AddComparisonClassify(info model.ComparisonClassify) rest.Result    //新增比较符分类
	ChangeComparisonClassify(info model.ComparisonClassify) rest.Result //根据id修改比较符分类信息
	DropComparisonClassify(ids []int64) rest.Result                     //删除比较符分类
	ObtainAllComparisonClassifyList() rest.Result                       //获取所有比较符分类
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainComparisonClassifyPage 根据条件获取比较符分类列表
func (s ServiceImpl) ObtainComparisonClassifyPage(mdl QueryPageMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items, total := s.repo.SelectPageComparisonClassifyByCondition(db, mdl)
	return rest.NewQueryPage(items, mdl.PageNumber, mdl.PageSize, total)
}

// ObtainComparisonClassifyInfo 获取单个比较符分类的信息
func (s ServiceImpl) ObtainComparisonClassifyInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	cInfo := s.repo.SelectComparisonClassifyById(db, id)
	if reflect.DeepEqual(cInfo, model.ComparisonClassify{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(cInfo, "ccInfo")
}

// AddComparisonClassify 新增比较符分类
func (s ServiceImpl) AddComparisonClassify(info model.ComparisonClassify) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertComparisonClassify(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeComparisonClassify 根据id修改比较符分类信息
func (s ServiceImpl) ChangeComparisonClassify(info model.ComparisonClassify) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateComparisonClassifyById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropComparisonClassify 删除比较符分类
func (s ServiceImpl) DropComparisonClassify(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteComparisonClassifyById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}

// ObtainAllComparisonClassifyList 获取所有比较符分类
func (s ServiceImpl) ObtainAllComparisonClassifyList() rest.Result {
	db := lib.ObtainCustomDb()
	comparisonClassifyList := s.repo.SelectAllComparisonClassifyList(db)
	return utils.Result(comparisonClassifyList, "comparisonClassifyList")
}
