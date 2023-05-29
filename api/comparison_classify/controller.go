package comparison_classify

import (
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/rest/req"
	"github.com/kataras/iris/v12"
)

type Controller struct {
	Ctx iris.Context
	Svc Service
}

func NewController(repository Repository) *Controller {
	return &Controller{Svc: NewService(repository)}
}

// GetComparisonClassifyPage
//
// @Summary 根据条件获取比较符分类信息分页列表
//
// @Description 根据条件获取比较符分类信息分页列表
//
// @Tags　比较符分类信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object QueryListMdl
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/classify/page [get]
func (c *Controller) GetComparisonClassifyPage() rest.Result {
	var queryMdl QueryPageMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainComparisonClassifyPage(queryMdl)
}

// GetComparisonClassifyBy
//
// @Summary 获取单个比较符分类信息的信息
//
// @Description 获取单个比较符分类信息信息
//
// @Tags　比较符分类信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param id path int true "比较符分类信息id"
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/classify/{id} [get]
func (c *Controller) GetComparisonClassifyBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择比较符分类信息", rest.ERROR)
	}
	return c.Svc.ObtainComparisonClassifyInfo(id)
}

// PostComparisonClassify
//
// @Summary 新增比较符分类信息
//
// @Description 新增比较符分类信息
//
// @Tags　比较符分类信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object model.ComparisonClassify
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/classify [post]
func (c *Controller) PostComparisonClassify() rest.Result {
	var processGroupInfo model.ComparisonClassify
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddComparisonClassify(processGroupInfo)
}

// PutComparisonClassify
//
// @Summary 根据id修改比较符分类信息信息
//
// @Description 根据id修改比较符分类信息信息
//
// @Tags　比较符分类信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object model.ComparisonClassify
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/classify [put]
func (c *Controller) PutComparisonClassify() rest.Result {
	var processGroupInfo model.ComparisonClassify
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeComparisonClassify(processGroupInfo)
}

// DeleteComparisonClassify
//
// @Summary 删除比较符分类信息
//
// @Description 删除比较符分类信息
//
// @Tags　比较符分类信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object req.DelModel
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/classify [delete]
func (c *Controller) DeleteComparisonClassify() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择比较符分类信息", rest.ERROR)
	}
	return c.Svc.DropComparisonClassify(delMdl.Ids)
}

// GetAllComparisonClassifyList
//
// @Summary 获取所有比较符分类
//
// @Description 获取所有比较符分类
//
// @Tags　比较符分类信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object req.DelModel
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/all/comparison/classify/list [delete]
func (c *Controller) GetAllComparisonClassifyList() rest.Result {
	return c.Svc.ObtainAllComparisonClassifyList()
}
