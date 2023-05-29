package subset

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

// GetSubList
//
// @Summary 根据条件获取控件子集列表
// @Description 根据条件获取控件子集列表
// @Tags　控件子集管理
// @Accept application/json
// @Produce application/json
// @Param object ListQueryMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/subset/sub/list [get]
func (c *Controller) GetSubList() rest.Result {
	var listQueryMdl ListQueryMdl
	if err := c.Ctx.ReadQuery(&listQueryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainSubList(listQueryMdl)
}

// GetSubBy
//
// @Summary 获取单个控件子集信息
// @Description 获取单个控件子集信息
// @Tags　控件子集管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "控件子集id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/subset/sub/{id} [get]
func (c *Controller) GetSubBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "参数错误", rest.ERROR)
	}
	return c.Svc.ObtainSingleSubInfo(id)
}

// PostSub
//
// @Summary 新增控件子集信息
// @Description 新增控件子集信息
// @Tags　控件子集管理
// @Accept application/json
// @Produce application/json
// @Param object SubDetail
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/subset/sub [post]
func (c *Controller) PostSub() rest.Result {
	var info model.Subset
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddSub(info)
}

// PutSub
//
// @Summary 修改控件子集信息
// @Description 修改控件子集信息
// @Tags　控件子集管理
// @Accept application/json
// @Produce application/json
// @Param object SubDetail
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/subset/sub [put]
func (c *Controller) PutSub() rest.Result {
	var info SubDetail
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if info.Id <= 0 {
		return rest.FailCustom(400, "id不能为空", rest.ERROR)
	}
	return c.Svc.ChangeSubById(info)
}

// DeleteSub
//
// @Summary 删除控件子集信息
// @Description 删除控件子集信息
// @Tags　控件子集管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Failure 400 {object} rest.Result
// @Router /flow/subset/sub [delete]
func (c *Controller) DeleteSub() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "ids不能为空", rest.ERROR)
	}
	return c.Svc.DropSub(delMdl.Ids)
}
