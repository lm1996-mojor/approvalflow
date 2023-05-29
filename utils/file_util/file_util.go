package file_util

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"five.com/lk_flow/api/flow_api/_const"
	"five.com/technical_center/core_library.git/log"
)

// ObjDataToJsonFile 对象数据转换成json文件
//
// @Parma path 文件保存路径
//
// @Parma filename 文件名称
//
// @Parma srcData 需要转换的数据
func ObjDataToJsonFile(path string, fileName string, srcData interface{}) error {
	err2 := checkParams(path, fileName)
	if err2 != nil {
		return err2
	}
	if srcData == nil {
		return errors.New("请指定数据对象")
	}
	// 创建文件(无论结果如何该文件夹都会存在)
	CreateDir(path)
	// 创建文件
	filePtr, err := os.Create(_const.APPROVAL_DATA_FILE_SAVE_PATH_PERFIX + _const.APPROVAL_DATA_FILE_SAVE_PATH_PERFIX + path + "/" + fileName + ".json")
	if err != nil {
		log.Error("文件创建错误")
		return err
	}
	defer filePtr.Close()
	//json化结构体实例
	data, err1 := json.MarshalIndent(srcData, "", "  ")
	if err1 != nil {
		log.Error("json文件数据解码错误: " + err1.Error())
		return err1
	}
	//写入文件
	filePtr.Write(data)

	return nil
}

// ReaderJsonFileToObj 读取json文件数据转换成指定对象
//
// @Parma path 文件保存路径
//
// @Parma filename 文件名称
//
// @Parma obj 承载数据的对象
func ReaderJsonFileToObj(path, fileName string, obj interface{}) (interface{}, error) {
	err2 := checkParams(path, fileName)
	if err2 != nil {
		return nil, err2
	}
	bytes, err := os.ReadFile(_const.APPROVAL_DATA_FILE_SAVE_PATH_PERFIX + path + "/" + fileName + ".json")
	if err != nil {
		return nil, err
	}
	//反向解码并给到实例
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// ChangeJsonFileData 更新指定json文件中的数据
//
// @Parma path 文件保存路径
//
// @Parma filename 文件名称
//
// @Parma obj 要更新的数据
func ChangeJsonFileData(path, fileName string, obj interface{}) (bool, error) {
	var mutex sync.Mutex
	err2 := checkParams(path, fileName)
	if err2 != nil {
		return false, err2
	}
	mutex.Lock()
	result, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		log.Error("文件转码错误：" + err.Error())
		return false, err
	}
	//写入到json文件中
	err = os.WriteFile(_const.APPROVAL_DATA_FILE_SAVE_PATH_PERFIX+path+"/"+fileName+".json", result, 0644)
	if err != nil {
		log.Error("文件写入错误：" + err.Error())
		return false, err
	}
	mutex.Unlock()
	return true, nil
}

// 参数检查
func checkParams(path, fileName string) error {
	if fileName == "" || len(fileName) <= 0 {
		return errors.New("请指定缓存的文件名")
	}
	if path == "" || len(path) <= 0 {
		return errors.New("请指定缓存文件的路径")
	}
	return nil
}

func RemoveJsonFile(path, fileName string) {
	err := os.Remove(_const.APPROVAL_DATA_FILE_SAVE_PATH_PERFIX + path + "/" + fileName + ".json")
	if err != nil {
		panic("删除json文件错误: " + err.Error())
	}
	err = os.Remove(path)
	if err != nil {
		panic("删除指定文件夹错误: " + err.Error())
	}
}
