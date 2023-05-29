package file_util

import (
	"os"

	"five.com/technical_center/core_library.git/log"
)

// HasDir 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// CreateDir 创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		log.Errorf("获取文件夹异常 -> %v\n", _err)
		panic("获取文件夹异常 -> %v\n" + _err.Error())
		return
	}
	if _exist {
		log.Infof("文件夹已存在！%s", path)
	} else {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Errorf("创建目录异常 -> %v\n", err)
			panic("创建目录异常 -> %v\n" + err.Error())
		}
	}
}
