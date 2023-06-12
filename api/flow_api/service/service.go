package service

import (
	"reflect"
	"strconv"

	"five.com/lk_flow/api/flow_api/_const"
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/engine"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/api/flow_api/repository/repoImpl"
	localUtils "five.com/lk_flow/utils"
	"five.com/lk_flow/utils/file_util"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/utils"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// ApprovalServiceService 业务逻辑接口
type ApprovalServiceService interface {
	InitiateApproval(params api_model.ApprovalParams) rest.Result               // 发起审批流程
	Approval(params api_model.CurrentApprovalStatus) rest.Result                // 参与人审批
	RevokeApprovalProcess(params api_model.CurrentApprovalStatus) rest.Result   // 撤回审批流程
	ObtainUserApprovalInfoPage(params api_model.QueryApprovalParam) rest.Result // 获取审批流程分页列表
	ObtainSingleApprovalInfoByApprovalCode(approvalCode string) rest.Result     // 根据审批编号获取单个审批信息
}

type ApprovalServiceImpl struct {
	repo repository.ProcessValueRepository
}

func NewService(repository repository.ProcessValueRepository) ApprovalServiceService {
	return &ApprovalServiceImpl{repo: repository}
}

var ctlValueRepo = repoImpl.CtlValueNewMysqlRepository()

// InitiateApproval 发起审批流程
func (s ApprovalServiceImpl) InitiateApproval(params api_model.ApprovalParams) rest.Result {
	tx := lib.ObtainCustomDbTx()
	//数据整理
	// 生成审批编号
	params.ApprovalCode = utils.GenerateCodeByUUID(5)
	// 获取租户id
	params.ClientId = localUtils.GetClientId()
	params.ProcessRate = 4
	for i := 0; i < len(params.CtlValues); i++ {
		params.CtlValues[i].ApprovalCode = params.ApprovalCode
	}
	for i := 0; i < len(params.PointDetails); i++ {
		// 设置所有节点的是否为当前审批节点的状态为：否
		params.PointDetails[i].IsCurrentPoint = false
		params.PointDetails[i].ApprovalCode = params.ApprovalCode
		params.PointDetails[i].PointRate = 5
	}
	// 新增审批信息
	approvalData := s.repo.InsertProcessValueInfo(tx, params)
	// 将数据转换为文件
	// 获取文件路径
	path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), params.ApprovalCode)
	// 创建审批信息json文件
	err := file_util.ObjDataToJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalData)
	if err != nil {
		log.Error("数据缓存文件创建失败： " + err.Error())
		panic(err)
	}
	// 启动流程
	engine.StartFlow(path)
	// 返回审批编号
	resultMap := make(map[string]interface{})
	resultMap["approvalCode"] = params.ApprovalCode
	return rest.SuccessResult(resultMap)
}

// Approval 参与人审批
func (s ApprovalServiceImpl) Approval(params api_model.CurrentApprovalStatus) rest.Result {
	/*
		1、获取指定的审批流程数据
		2、修改前如果该节点审批形式为 “或签” 则需要查询该节点是否已经进行了审批
		3、根据节点审批形式修改参与人的审批结果，且审批过程中应加锁操作
		4、根据节点中参与人的审批情况修改节点审批进度
		5、根据节点审批进度修改流程的审批进度
	*/
	params.ClientId = localUtils.GetClientId()
	path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), params.ApprovalCode)
	approvalParams := getApprovalParams(path)
	var err error
	// 事务数据
	txApprovalParams := approvalParams
	if approvalParams.ProcessRate < 4 {
		return rest.FailCustom(400, "流程已结束，请勿重复操作", rest.ERROR)
	}
	if params.ExamineType == 2 {
		isApproval := engine.CheckPointIsApproval(approvalParams.PointDetails, params.NodeDetailId)
		if isApproval {
			return rest.FailCustom(400, "该节点审批形式为或签，已经有人审批了。无需再次审批", rest.ERROR)
		}
	}

	// 参与人审批
endloop:
	for i := 0; i < len(approvalParams.PointDetails); i++ {
		if approvalParams.PointDetails[i].Id == params.NodeDetailId {
			for j := 0; j < len(approvalParams.PointDetails[i].ParticipantInfos); j++ {
				if approvalParams.PointDetails[i].ParticipantInfos[j].ObjId == params.Participant.ObjId {
					// 参与人审批核心
					// 参与人进行审批
					approvalParams.PointDetails[i].ParticipantInfos = engine.ParticipantApproval(approvalParams.PointDetails[i].ParticipantInfos, params.Participant)
					/*
							审批结果: （1 同意 2退回 3驳回 4审批中 5待执行 6无操作 7 撤销）
								同意： 进入下一个阶段
								退回： 退回到上一个审批节点，并且要有退回意见,且流程不停止，但不进入下一个节点，修改包含退回到的节点与当前节点之间的所有节点进度及参与人的审批结果，并将数据保存至缓存文件中
								驳回： 即不同意，流程结束。流程进度为驳回，当前节点进度为驳回。参与人审批结果为驳回，且必须要有驳回意见。其他数据保持不变。
						     		  并将数据保存至数据库中
								其他状态为未操作状态，无需处理
					*/
					switch params.Participant.ApprovalResult {
					case 2: //退回： 退回到上一个审批节点，并且要有退回意见,且流程不停止，但不进入下一个节点，修改包含退回到的节点与当前节点之间的所有节点进度及参与人的审批结果，并将数据保存至缓存文件中
						checkPointList := make([]api_model.PointDetail, 0)
						ids := make([]int64, 0)
						var backId int64
						currentPointValue := engine.ObtainCurrentPointValue(approvalParams.PointDetails)
						if currentPointValue.Id != params.NodeDetailId {
							return rest.FailCustom(500, "请确认当前的节点是否符合规范", rest.ERROR)
						}
						// 判别是退回到上一个审批节点，还是退回到其他节点
						if params.BackNodeId != 0 { // 退回至其他节点
							for _, id := range params.BackIntervalNodeIds {
								checkPointList = append(checkPointList, engine.ObtainPointValueInfo(approvalParams.PointDetails, id))
							}
							result := checkApprovalPointDataConformanceToSpecifications(checkPointList, true)
							if !reflect.DeepEqual(result, rest.Result{}) {
								return result
							}
							ids = append(ids, params.BackIntervalNodeIds...)
							ids = append(ids, params.NodeDetailId)
							ids = append(ids, params.BackNodeId)
							backId = params.BackNodeId
						} else { // 退回到上一个审批节点
							ids = append(ids, params.NodeDetailId)
							ids = append(ids, engine.FindThePreviousApprovalNodeOfTheSpecifiedNode(params.NodeDetailId, approvalParams.PointDetails)...)
							backId = engine.FindPreviousApprovalId(params.NodeDetailId, approvalParams.PointDetails)
						}
						checkPointList = make([]api_model.PointDetail, 0)
						checkPointList = append(checkPointList, engine.ObtainPointValueInfo(approvalParams.PointDetails, backId))
						result := checkApprovalPointDataConformanceToSpecifications(checkPointList, false)
						if !reflect.DeepEqual(result, rest.Result{}) {
							return result
						}
						// 重置数据状态
						approvalParams.PointDetails = engine.ResetBatchPointApprovalStatus(approvalParams.PointDetails, ids)
						approvalParams.PointDetails = engine.ResetSinglePointApprovalStatus(approvalParams.PointDetails, backId)
						// 更新流程数据
						_, err = file_util.ChangeJsonFileData(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalParams)
						if err != nil {
							log.Error("更新流程数据失败")
							panic(err)
						}
						return rest.SuccessCustom("退回成功", nil, rest.Success)
					case 3: //驳回： 即不同意，流程结束。不进入下一个步骤,流程进度为驳回，当前节点进度为驳回。参与人审批结果为驳回，且必须要有驳回意见。其他数据保持不变。并将数据保存至数据库中
						approvalParams.ProcessRate = 3
						for k := 0; k < len(approvalParams.PointDetails); k++ {
							if approvalParams.PointDetails[k].Id == params.NodeDetailId {
								approvalParams.PointDetails[k].PointRate = 3
								approvalParams.PointDetails[k].ParticipantInfos = engine.ParticipantApproval(approvalParams.PointDetails[k].ParticipantInfos, params.Participant)
							}
						}
						file_util.RemoveJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME)
						// TODO: 发送消息（流程结束，所有参与人）
						// 更新数据库流程相关数据
						resultStatus := s.renewProcessData(approvalParams)
						if !resultStatus { // 结果取反,如果数据更新失败则进行文件数据回滚
							err = file_util.ObjDataToJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, txApprovalParams)
							if err != nil {
								log.Error("事务数据文件回滚失败")
								panic("服务器错误")
							}
						}
						return rest.SuccessCustom("驳回成功,流程已结束", nil, rest.Success)
					}
					// 跳出循环到指定位置
					break endloop
				}
			}
		}
	}

	// 节点处理核心
	for i := 0; i < len(approvalParams.PointDetails); i++ {
		if approvalParams.PointDetails[i].Id == params.NodeDetailId {
			// 设置是否检测全部（即代表是否为会签/或签）
			// false 代表或签,即不检测全部，仅需有一个审批过即可
			isAll := false
			flag := false

			//ExamineType		|	审批形式				|（1 会签 2 或签）
			switch params.ExamineType {
			case 1:
				// 代表会签
				isAll = true
			}

			// 判断是否能够进入下一个节点
			if engine.CheckParticipantsApprovalResult(approvalParams.PointDetails[i].ParticipantInfos, 1, isAll) {
				flag = true
			}
			// 判断是否能够进入下一个节点
			if flag {

				// 判断下一个节点是否为结束节点
				approvalParams.PointDetails[i].PointRate = 1
				if approvalParams.PointDetails[i].NextStepType == 6 {
					// 已经到了结束节点
					approvalParams.ProcessRate = 1
					file_util.RemoveJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME)
					// TODO: 发送消息（流程结束,所有参与人）
					// 更新数据库流程相关数据
					resultStatus := s.renewProcessData(approvalParams)
					if !resultStatus { // 结果取反,如果数据更新失败则进行文件数据回滚
						err = file_util.ObjDataToJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, txApprovalParams)
						if err != nil {
							log.Error("事务数据文件回滚失败")
							panic("服务器错误")
						}
					}
					return rest.SuccessCustom("审批成功，流程已结束", nil, rest.Success)
				} else { // 如果下一个不为结束节点，即进行下一个节点审批
					// 判定是否为结束节点位置
					approvalParams.PointDetails = engine.NextStepPoint(approvalParams.PointDetails, approvalParams.PointDetails[i])
					currentNodeInfo := engine.ObtainCurrentPointValue(approvalParams.PointDetails)
					// 判断是否为结束位置
					if currentNodeInfo.PointType == 6 {
						approvalParams.ProcessRate = 1
						resultStatus := s.renewProcessData(approvalParams)
						if !resultStatus { // 结果取反,如果数据更新失败则进行文件数据回滚
							err = file_util.ObjDataToJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, txApprovalParams)
							if err != nil {
								log.Error("事务数据文件回滚失败")
								panic("服务器错误")
							}
						}
						file_util.RemoveJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME)
						// TODO: 发送消息（流程结束,所有参与人）
						// 更新数据库流程相关数据
						return rest.SuccessCustom("审批成功，流程已结束", nil, rest.Success)
					} else {
						// 更新流程数据
						_, err = file_util.ChangeJsonFileData(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalParams)
						if err != nil {
							log.Error("更新流程数据失败")
							panic(err)
						}
					}
				}
			} else {
				// 更新流程数据
				_, err = file_util.ChangeJsonFileData(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalParams)
				if err != nil {
					log.Error("更新流程数据失败")
					panic(err)
				}
			}
			break
		}
	}
	return rest.SuccessCustom("审批成功", nil, rest.Success)
}

// 更新流程数据
func (s ApprovalServiceImpl) renewProcessData(approvalParams api_model.ApprovalParams) bool {
	tx := lib.ObtainCustomDbTx()
	return s.repo.UpdateProcessValueInfoByApprovalCode(tx, approvalParams)
}

// RevokeApprovalProcess 撤回审批流程
func (s ApprovalServiceImpl) RevokeApprovalProcess(params api_model.CurrentApprovalStatus) rest.Result {
	/*
		撤回审批流程思路
		1、在撤回前需要检查除发起人外流程是否已经有人审批过了，如果没有则可以进行撤回操作，如果有则提示该流程已有人审批，无法撤回
		2、撤回流程需要删除缓存文件，同时软删数据库中的数据
	*/
	path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), params.ApprovalCode)
	approvalParams := getApprovalParams(path)
	//txApprovalParams := approvalParams
	for _, detail := range approvalParams.PointDetails {
		if detail.PointType == 5 {
			for _, pointDetail := range approvalParams.PointDetails {
				if pointDetail.Id == detail.NextStep {
					if engine.CheckWhetherTheParticipantHasBeenApproved(pointDetail.ParticipantInfos) {
						return rest.FailCustom(400, "该流程已有人审批，无法撤回", rest.ERROR)
					}
				}
			}
		}
	}
	//if !s.repo.DeleteProcessValueInfo(lib.ObtainCustomDbTx(), approvalParams) {
	//	log.Error("删除流程信息失败")
	//	panic("服务器错误")
	//}
	// 修改流程进度为撤销状态
	approvalParams.ProcessRate = 6
	approvalParams.PointDetails = engine.DropAllNodesExceptTheInitiatorNode(approvalParams.PointDetails)
	approvalParams.PointDetails[0].PointRate = 6
	approvalParams.PointDetails[0].ParticipantInfos[0].ApprovalResult = 7
	approvalParams.PointDetails[0].ParticipantInfos[0].Opinions = "流程已撤销"
	s.repo.UpdateProcessValueInfoByApprovalCode(lib.ObtainCustomDbTx(), approvalParams)
	file_util.RemoveJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME)
	return rest.SuccessCustom("撤销成功", nil, rest.Success)
}

// ObtainUserApprovalInfoPage 获取审批流程分页列表
func (s ApprovalServiceImpl) ObtainUserApprovalInfoPage(params api_model.QueryApprovalParam) rest.Result {
	params.ClientId = localUtils.GetClientId()
	db := lib.ObtainCustomDb()

	var currentApprovalParams []api_model.RepoApprovalParam
	var total int64
	// 查询类型（1 待处理 2 已处理 3 已发起 4 收到的）
	switch params.QueryType {
	case 1: // 待处理
		// 查询出该审批流程中的所在节点的进度为审批中
		// 获取该人员参与的审批流程（审批人/抄送人）
		items, totalNumber := s.repo.SelectApprovalOfUserParticipation(db, params)
		approvalCodeList := make([]string, 0)
		for _, item := range items {
			path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
			item = getApprovalParams(path)
			if engine.CheckCurrentApprovalNodeIsContainUser(item.PointDetails, params.UserId) {
				approvalCodeList = append(approvalCodeList, item.ApprovalCode)
			}
		}
		items, totalNumber = s.repo.SelectProcessValueListByApprovalList(lib.ObtainCustomDb(), approvalCodeList, params)
		for _, item := range items {
			path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
			item = getApprovalParams(path)
			currentNodeInfo := engine.ObtainCurrentPointValue(item.PointDetails)
			currentApprovalParams = append(currentApprovalParams, assignmentRepoApprovalParam(item, currentNodeInfo))
		}
		total = totalNumber
	case 2: // 已处理
		// 查询该用户在审批流程中的已经审批过的
		items, totalNumber := s.repo.SelectApprovalOfUserParticipation(db, params)
		approvalCodeList := make([]string, 0)
		for _, item := range items {
			if item.ProcessRate == 1 || item.ProcessRate == 3 || item.ProcessRate == 6 {
				item = s.repo.SelectSingleProcessValue(lib.ObtainCustomDb(), item.ApprovalCode)
			} else {
				// 获取该流程的状态和数据
				path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
				item = getApprovalParams(path)
			}
			if engine.CheckIfThereAreAnyProcessedItemsForTheUserInTheApproval(item.PointDetails, params.UserId) {
				approvalCodeList = append(approvalCodeList, item.ApprovalCode)
			}
		}
		items, totalNumber = s.repo.SelectProcessValueListByApprovalList(lib.ObtainCustomDb(), approvalCodeList, params)
		for _, item := range items {
			path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
			item = getApprovalParams(path)
			currentNodeInfo := engine.ObtainCurrentPointValue(item.PointDetails)
			currentApprovalParams = append(currentApprovalParams, assignmentRepoApprovalParam(item, currentNodeInfo))
		}
		total = totalNumber
	case 3: // 已发起
		// 查询出发起人属于该节点的
		items, totalNumber := s.repo.SelectInitiatedApprovalByUserId(db, params)
		for _, item := range items {
			// 进度为：同意、驳回、撤销，则去数据库中查询数据，其他的则去获取json文件中的数据
			if item.ProcessRate == 1 || item.ProcessRate == 3 || item.ProcessRate == 6 {
				item = s.repo.SelectSingleProcessValue(lib.ObtainCustomDb(), item.ApprovalCode)
			} else {
				// 获取该流程的状态和数据
				path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
				item = getApprovalParams(path)
			}
			currentNodeInfo := engine.ObtainCurrentPointValue(item.PointDetails)
			currentApprovalParams = append(currentApprovalParams, assignmentRepoApprovalParam(item, currentNodeInfo))
		}
		total = totalNumber
	case 4: // 收到的
		// 查询抄送人节点中是否存在该用户，且已经通过的
		items, totalNumber := s.repo.SelectApprovalOfUserParticipation(db, params)
		approvalCodeList := make([]string, 0)
		for _, item := range items {
			if item.ProcessRate == 1 || item.ProcessRate == 3 || item.ProcessRate == 6 {
				item = s.repo.SelectSingleProcessValue(lib.ObtainCustomDb(), item.ApprovalCode)
			} else {
				// 获取该流程的状态和数据
				path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
				item = getApprovalParams(path)
			}
			if engine.CheckIfThereAreAnyReceivedByTheUserInTheApproval(item.PointDetails, params.UserId) {
				approvalCodeList = append(approvalCodeList, item.ApprovalCode)
			}
		}
		items, totalNumber = s.repo.SelectProcessValueListByApprovalList(lib.ObtainCustomDb(), approvalCodeList, params)
		for _, item := range items {
			path := getCacheFilePath(params.AppCode, strconv.FormatInt(params.ClientId, 10), item.ApprovalCode)
			item = getApprovalParams(path)
			currentNodeInfo := engine.ObtainCurrentPointValue(item.PointDetails)
			currentApprovalParams = append(currentApprovalParams, assignmentRepoApprovalParam(item, currentNodeInfo))
		}
		total = totalNumber
	default: // 返回错误信息
		return rest.FailCustom(400, "无效的查询类型", rest.ERROR)
	}
	return rest.NewQueryPage(currentApprovalParams, params.PageNumber, params.PageSize, total)
}

// ObtainSingleApprovalInfoByApprovalCode 根据审批编号获取单个审批信息
func (s ApprovalServiceImpl) ObtainSingleApprovalInfoByApprovalCode(approvalCode string) rest.Result {
	db := lib.ObtainCustomDb()
	approvalParams := s.repo.SelectSingleProcessValue(db, approvalCode)
	// 判断该审批流程是否已经完成审批,如果没有完成则该审批流程信息需要去数据库中获取
	// 流程结果进度（1 同意 2退回 3驳回 4审批中 5待执行 6撤销）
	if approvalParams.ProcessRate == 2 || approvalParams.ProcessRate == 4 || approvalParams.ProcessRate == 5 {
		path := getCacheFilePath(approvalParams.AppCode, strconv.FormatInt(approvalParams.ClientId, 10), approvalParams.ApprovalCode)
		approvalParams = getApprovalParams(path)
		// 查询该审批中的表单数据
		approvalParams.CtlValues = ctlValueRepo.SelectCtlValueListByApprovalCode(db, approvalCode)
	}
	if reflect.DeepEqual(approvalParams, api_model.ApprovalParams{}) {
		return rest.FailCustom(500, "该审批流程不存在", rest.ERROR)
	}
	// 获取当前审批流程的节点进度
	currentNodeInfo := engine.ObtainCurrentPointValue(approvalParams.PointDetails)
	var repoSingleApprovalParam api_model.RepoSingleApprovalParam
	// 拼装响应数据
	repoSingleApprovalParam.ApprovalInfo = approvalParams
	repoSingleApprovalParam.CurrentApprovalInfo = assignmentRepoApprovalParam(approvalParams, currentNodeInfo)
	resultMap := make(map[string]interface{})
	resultMap["repoSingleApprovalParam"] = repoSingleApprovalParam
	return rest.SuccessCustom("查询成功", resultMap, rest.Success)
}

// 赋值数据给响应对象(api_model.RepoApprovalParam)
func assignmentRepoApprovalParam(approvalParams api_model.ApprovalParams, currentNodeInfo api_model.PointDetail) (currentApprovalParamInfo api_model.RepoApprovalParam) {
	currentApprovalParamInfo.ApprovalId = approvalParams.Id
	currentApprovalParamInfo.ProcessId = approvalParams.ProcessId
	currentApprovalParamInfo.ApprovalCode = approvalParams.ApprovalCode
	currentApprovalParamInfo.ApprovalTitle = approvalParams.ApprovalTitle
	currentApprovalParamInfo.CtlValues = ctlValueRepo.SelectCtlValueListByApprovalCode(lib.ObtainCustomDb(), approvalParams.ApprovalCode)
	currentApprovalParamInfo.Participant = currentNodeInfo.ParticipantInfos
	currentApprovalParamInfo.NodeDetailId = currentNodeInfo.Id
	currentApprovalParamInfo.ExamineType = currentNodeInfo.ExamineType
	currentApprovalParamInfo.PointType = currentNodeInfo.PointType
	currentApprovalParamInfo.PointRate = currentNodeInfo.PointRate
	currentApprovalResult := ""
	if approvalParams.ProcessRate == 6 {
		currentApprovalResult = "流程已撤销"
		currentApprovalParamInfo.ApprovalIsEnd = true
	} else if len(currentNodeInfo.ParticipantInfos) > 1 {
		examineTypeStr := ""
		if currentNodeInfo.ExamineType == 1 {
			examineTypeStr = "会签"
		} else {
			examineTypeStr = "或签"
		}
		currentApprovalResult = examineTypeStr + "审批中"
	} else if len(currentNodeInfo.ParticipantInfos) == 1 {
		if currentNodeInfo.PointRate == 3 {
			for _, info := range currentNodeInfo.ParticipantInfos {
				if info.ApprovalResult == 3 {
					currentApprovalResult = info.ObjName + "已驳回"
					currentApprovalParamInfo.ApprovalIsEnd = true
					break
				}
			}
		} else {
			currentApprovalResult = currentNodeInfo.ParticipantInfos[0].ObjName + "审批中"
		}
	} else if currentNodeInfo.PointType == 6 {
		currentApprovalResult = "审批通过,流程已结束"
		currentApprovalParamInfo.ApprovalIsEnd = true
	} else if currentNodeInfo.PointRate == 3 {
		for _, info := range currentNodeInfo.ParticipantInfos {
			if info.ApprovalResult == 3 {
				currentApprovalResult = info.ObjName + "已驳回"
				currentApprovalParamInfo.ApprovalIsEnd = true
				break
			}
		}
	} else {
		log.Error("审批结果拼装失败,该节点没有参与人")
		panic("服务器错误")
	}
	currentApprovalParamInfo.CurrentApprovalResult = currentApprovalResult
	return currentApprovalParamInfo
}

// 判断指定流程是否已经结束
func isApprovalEnd(approvalParams api_model.ApprovalParams) bool {
	if approvalParams.ProcessRate < 4 {
		return true
	}
	return false
}

// 获取缓存文件路径
func getCacheFilePath(appCode, clientId string, approvalCode string) string {
	return appCode + "/" + clientId + "/" + approvalCode
}

// 获取文件路径
func getApprovalParams(path string) (approvalParams api_model.ApprovalParams) {
	// 读取文件数据
	obj, err := file_util.ReaderJsonFileToObj(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, api_model.ApprovalParams{})
	if err != nil {
		log.Error("json文件读取转化失败： " + err.Error())
		panic(err)
	}
	// 断言类型
	approvalParams = obj
	return approvalParams
}

// 检查审批节点的数据是否符合规范
// isIncludeCCRecipient：是否需要包含抄送节点检查
func checkApprovalPointDataConformanceToSpecifications(pointList []api_model.PointDetail, isIncludeCCRecipient bool) rest.Result {
	for _, pointValueInfo := range pointList {
		if reflect.DeepEqual(pointValueInfo, api_model.PointDetail{}) {
			return rest.FailCustom(500, "退回的节点不规范", rest.ERROR)
		}
		if isIncludeCCRecipient {
			// 包含抄送节点检查
			if (pointValueInfo.PointType != 1 && pointValueInfo.PointType != 2) || ((pointValueInfo.PointType == 2 || pointValueInfo.PointType == 1) && pointValueInfo.PointRate != 1) {
				return rest.FailCustom(500, "退回的节点不规范", rest.ERROR)
			}
		} else { // 不包含抄送节点检查
			if (pointValueInfo.PointType != 1) || (pointValueInfo.PointType == 1 && pointValueInfo.PointRate != 1) {
				return rest.FailCustom(500, "退回的节点不规范", rest.ERROR)
			}
		}
	}
	return rest.Result{}
}
