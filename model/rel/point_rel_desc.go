package rel

// PointRelDesc 流程节点关系描述表
type PointRelDesc struct {
	Id               int64 `json:"id,omitempty"`               // 主键id
	PointId          int64 `json:"pointId,omitempty"`          // 节点id
	PreviousStepType int8  `json:"previousStepType,omitempty"` // 上一节点类型(1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点)
	PreviousStep     int64 `json:"previousStep,omitempty"`     // 上一节点
	NextStep         int64 `json:"nextStep,omitempty"`         // 下一节点
	NextStepType     int8  `json:"nextStepType,omitempty"`     // 下一节点类型1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点
}

// 定义列
func (c *PointRelDesc) allColumn() []string {
	columns := []string{"id", "point_id", "previous_step_type", "previous_step", "next_step", "next_step_type"}
	return columns
}

// GetAllColumn 获取所有列
func (c *PointRelDesc) GetAllColumn() (result string) {
	columns := c.allColumn()
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			result += columns[i]
		} else {
			result += columns[i] + " , "
		}
	}
	return result
}

// TableName 获取表名
func (c *PointRelDesc) TableName() string {
	return "point_rel_desc"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *PointRelDesc) GetAllColumWithAlias(alias string) (result string) {
	columns := c.allColumn()
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			result += alias + "." + columns[i]
		} else {
			result += alias + "." + columns[i] + " , "
		}
	}
	return result
}
