package participant_format

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

// GetParticipantFormatPage
//
// @Summary 根据条件获取参与者形式列表
// @Description 根据条件获取参与者形式列表
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/participant/format/page [get]
func (c *Controller) GetParticipantFormatPage() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainParticipantFormatPage(queryMdl)
}

// GetParticipantFormatBy
//
// @Summary 获取单个参与者形式的信息
// @Description 获取单个参与者形式信息
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "参与者形式id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/participant/format/{id} [get]
func (c *Controller) GetParticipantFormatBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择参与者形式", rest.ERROR)
	}
	return c.Svc.ObtainParticipantFormatInfo(id)
}

// PostParticipantFormat
//
// @Summary 新增参与者形式的信息
// @Description 新增参与者形式信息
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Param object model.ParticipantFormat
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/participant/format [get]
func (c *Controller) PostParticipantFormat() rest.Result {
	var processInfo model.ParticipantFormat
	if err := c.Ctx.ReadJSON(&processInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddParticipantFormat(processInfo)
}

// PutParticipantFormat
//
// @Summary 根据id修改参与者形式信息
// @Description 根据id修改参与者形式信息
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Param object model.ParticipantFormat
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/participant/format [get]
func (c *Controller) PutParticipantFormat() rest.Result {
	var processInfo model.ParticipantFormat
	if err := c.Ctx.ReadJSON(&processInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeParticipantFormat(processInfo)
}

// DeleteParticipantFormat
//
// @Summary 删除参与者形式
// @Description 删除参与者形式
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/participant/format [get]
func (c *Controller) DeleteParticipantFormat() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择参与者形式", rest.ERROR)
	}
	return c.Svc.DropParticipantFormat(delMdl.Ids)
}

// GetParticipantFormatAllList
//
// @Summary 获取全部参与者形式
// @Description 获取全部参与者形式
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/participant/format/page [get]
func (c *Controller) GetParticipantFormatAllList() rest.Result {
	return c.Svc.ObtainParticipantFormatAllList()
}
