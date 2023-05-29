package engine

import (
	"five.com/lk_flow/api/flow_api/_const"
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/utils/file_util"
	"five.com/technical_center/core_library.git/utils/trans"
)

/*
	通用信息：
	|————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————|
	|	使用字段  		|	字段描述			    |  字段包含值																|		包含值解释																		 |
	|———————————————————|———————————————————————|———————————————————————————————————————————————————————————————————————|————————————————————————————————————————————————————————————————————————————————————————|
	| ProcessRate		|	流程审批进度   	    |（1 同意 2退回 3驳回 4审批中 5待执行）										|		/																				 |
	| PointRate			|	节点审批进度    	    |（1 同意 2退回 3驳回 4审批中 5待执行）										|		/																				 |
	| ApprovalResult	|	参与人审批状态  	   	|（1 同意 2退回 3驳回 4审批中 5待执行 6无操作）								|		/																				 |
	| PointType			|	节点类型				|（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）		|		/																				 |
	| ExamineType		|	审批形式				|（1 会签 2 或签） 														|		1  会签		所有参与人都开始审核，单必须是所有的参与人审批都通过才可以进行，节点进度修改		 |
	|					|						|																		|		2  或签		任意一个参与人审批通过，则修改节点的进度为通过								 |
	|					|						|																		|																						 |
	|————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————|
*/

// StartFlow 启动流程
func StartFlow(path string) {
	// 开始执行流程
	objData, err := file_util.ReaderJsonFileToObj(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, api_model.ApprovalParams{})
	if err != nil {
		panic("获取缓存数据失败: " + err.Error())
	}
	approvalParams := objData.(api_model.ApprovalParams)
	// 发起人节点对象
	var promoterPointDetail api_model.PointDetail
	// 发起人数据处理
	for _, pointDetail := range approvalParams.PointDetails {
		if pointDetail.PointType == 5 {
			trans.DeepCopy(pointDetail, &promoterPointDetail)
		}
	}
	promoterPointDetail.PointRate = 1
	approvalParams.PointDetails = ApprovalPointChangeFun(approvalParams.PointDetails, promoterPointDetail, approvalParams)
	// 流程进入到下一个审批节点
	approvalParams.PointDetails = NextStepPoint(approvalParams.PointDetails, promoterPointDetail, approvalParams)
	ok, err2 := file_util.ChangeJsonFileData(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalParams)
	if !ok {
		file_util.RemoveJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME)
		panic("流程启动失败： " + err2.Error())
	}
}
