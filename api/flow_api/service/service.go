package service

import (
	"strconv"

	"five.com/lk_flow/api/flow_api/_const"
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/engine"
	"five.com/lk_flow/api/flow_api/repository"
	localUtils "five.com/lk_flow/utils"
	"five.com/lk_flow/utils/file_util"
	"five.com/technical_center/core_library.git/log"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/utils"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// ApprovalServiceService 业务逻辑接口
type ApprovalServiceService interface {
	InitiateApproval(params api_model.ApprovalParams) rest.Result             // 发起审批流程
	Approval(params api_model.CurrentApprovalStatus) rest.Result              // 参与人审批
	RevokeApprovalProcess(params api_model.CurrentApprovalStatus) rest.Result // 撤回审批流程
	FindUserApprovalInfo(params api_model.QueryApprovalParam) rest.Result
}

type ApprovalServiceImpl struct {
	repo repository.ProcessValueRepository
}

func NewService(repository repository.ProcessValueRepository) ApprovalServiceService {
	return &ApprovalServiceImpl{repo: repository}
}

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
		params.PointDetails[i].ApprovalCode = params.ApprovalCode
		params.PointDetails[i].PointRate = 5
	}
	// 新增审批信息
	approvalData := s.repo.InsertProcessValueInfo(tx, params)
	// 将数据转换为文件
	path := params.AppCode + "/" + params.BusinessCode + "/" + strconv.FormatInt(params.ClientId, 10)
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
	path := getCacheFilePath(params.AppCode, params.BusinessCode, strconv.FormatInt(params.ClientId, 10))
	approvalParams := getApprovalParams(path)
	var err error
	// 事务数据
	txApprovalParams := approvalParams
	if approvalParams.ProcessRate > 3 {
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
							审批结果: 1 同意 2退回 3驳回 4审批中 5待执行 6无操作
								同意： 进入下一个阶段
								退回： 退回到上一个审批节点，并且要有退回意见,且流程不停止，但不进入下一个节点，修改包含退回到的节点与当前节点之间的所有节点进度及参与人的审批结果，并将数据保存至缓存文件中
								驳回： 即不同意，流程结束。流程进度为驳回，当前节点进度为驳回。参与人审批结果为驳回，且必须要有驳回意见。其他数据保持不变。
						     		  并将数据保存至数据库中
								其他状态为未操作状态，无需处理
					*/
					switch params.Participant.ApprovalResult {
					case 2: //退回： 退回到上一个审批节点，并且要有退回意见,且流程不停止，但不进入下一个节点，修改包含退回到的节点与当前节点之间的所有节点进度及参与人的审批结果，并将数据保存至缓存文件中
						ids := make([]int64, 0)
						// 判别是退回到上一个审批节点，还是退回到其他节点
						if params.BackNodeId != 0 { // 退回至其他节点
							for _, detail := range approvalParams.PointDetails {
								if params.BackNodeId == detail.Id {
									ids = append(ids, detail.Id)
									ids = append(ids, engine.FindNeedChangeNodeIds(detail.Id, params.NodeDetailId, detail.NextStep, approvalParams.PointDetails)...)
									break
								}
							}
						} else { // 退回到上一个审批节点
							ids = append(ids, params.NodeDetailId)
							ids = append(ids, engine.FindThePreviousApprovalNodeOfTheSpecifiedNode(params.NodeDetailId, approvalParams.PointDetails)...)
						}
						// 重置数据状态
						approvalParams.PointDetails = engine.ResetPointApprovalStatus(approvalParams.PointDetails, ids)
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
				if approvalParams.PointDetails[i].NextStepType == 6 {
					// 已经到了结束节点
					approvalParams.PointDetails[i].PointRate = 1
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
					approvalParams.PointDetails = engine.NextStepPoint(approvalParams.PointDetails, approvalParams.PointDetails[i], approvalParams)
					// 更新流程数据
					_, err = file_util.ChangeJsonFileData(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalParams)
					if err != nil {
						log.Error("更新流程数据失败")
						panic(err)
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
	path := getCacheFilePath(params.AppCode, params.BusinessCode, strconv.FormatInt(params.ClientId, 10))
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
	if !s.repo.DeleteProcessValueInfo(lib.ObtainCustomDbTx(), approvalParams) {
		log.Error("删除流程信息失败")
		panic("服务器错误")
	}
	file_util.RemoveJsonFile(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME)
	return rest.SuccessCustom("撤销成功", nil, rest.Success)
}

func (s ApprovalServiceImpl) FindUserApprovalInfo(params api_model.QueryApprovalParam) rest.Result {
	//TODO implement me
	panic("implement me")
}

// 获取缓存文件路径
func getCacheFilePath(appCode, businessCode, clientId string) string {
	return appCode + "/" + businessCode + "/" + clientId
}

// 获取文件路径
func getApprovalParams(path string) (approvalParams api_model.ApprovalParams) {
	// 读取文件数据
	obj, err := file_util.ReaderJsonFileToObj(path, _const.APPROVAL_DATA_FILE_SAVE_FILENAME, approvalParams)
	if err != nil {
		log.Error("json文件读取转化失败： " + err.Error())
		panic(err)
	}
	// 断言类型
	approvalParams = obj.(api_model.ApprovalParams)
	return approvalParams
}
