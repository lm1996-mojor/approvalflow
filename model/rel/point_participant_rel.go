package rel

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type PointParticipantRel struct {
	Id                  int64 `json:"id,omitempty"`                  // 主键id
	PointId             int64 `json:"pointId,omitempty"`             // 节点id
	ParticipantFormatId int64 `json:"participantFormatId,omitempty"` // 参与者形式id
}

// 定义列
func (c *PointParticipantRel) allColumn() []string {
	columns := []string{"id", "point_id", "participant_format_id"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *PointParticipantRel) GetAllColumn() (result string) {
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
func (c *PointParticipantRel) TableName() string {
	return "point_participant_rel"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *PointParticipantRel) GetAllColumWithAlias(alias string) (result string) {
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
