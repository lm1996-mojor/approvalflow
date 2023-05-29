package control_info

import (
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest"
	"github.com/kataras/iris/v12"
)

type Controller struct {
	Ctx iris.Context
	Svc Service
}

func NewController(repository Repository) *Controller {
	return &Controller{Svc: NewService(repository)}
}

// GetCtlList
//
// @Summary 根据条件获取控件列表
// @Description 根据条件获取控件列表
// @Tags　控件管理
// @Accept application/json
// @Produce application/json
// @Param object ListQueryMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/control/ctl/list [get]
func (c *Controller) GetCtlList() rest.Result {
	var listQueryMdl ListQueryMdl
	if err := c.Ctx.ReadQuery(&listQueryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainCtlList(listQueryMdl)
}

// GetCtlBy
//
// @Summary 获取单个控件信息
// @Description 获取单个控件信息
// @Tags　控件管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "控件id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/control/ctl/{id} [get]
func (c *Controller) GetCtlBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "参数错误", rest.ERROR)
	}
	return c.Svc.ObtainSingleCtlInfo(id)
}

// PostCtl
//
// @Summary 新增控件信息
// @Description 新增控件信息
// @Tags　控件管理
// @Accept application/json
// @Produce application/json
// @Param object InformationControl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/control/ctl [post]
func (c *Controller) PostCtl() rest.Result {
	var infos []model.ControlInfo
	if err := c.Ctx.ReadJSON(&infos); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if infos[0].ParentId <= 0 {
		return rest.FailCustom(400, "请选择具体流程", rest.ERROR)
	}
	for i := 0; i < len(infos); i++ {
		if infos[i].OwnerType <= 0 {
			infos[i].OwnerType = 1
		}
	}

	return c.Svc.AddCtl(infos)
}

// PutCtl
//
// @Summary 修改控件信息
// @Description 修改控件信息
// @Tags　控件管理
// @Accept application/json
// @Produce application/json
// @Param object InformationControl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/control/ctl [put]
func (c *Controller) PutCtl() rest.Result {
	var info model.ControlInfo
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if info.Id <= 0 {
		return rest.FailCustom(400, "id不能为空", rest.ERROR)
	}
	return c.Svc.ChangeCtlById(info)
}
