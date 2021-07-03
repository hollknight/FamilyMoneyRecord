package resource_utils

import (
	"FamilyMoneyRecord/config"
	"io/ioutil"
	"os"
)

// IsExist 判断该路径下该文件是否存在
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DeleteFile 删除指定文件
func DeleteFile(path string) error {
	err := os.Remove(path)
	return err
}

// GetJSONName 获取json文件夹下所有json文件名
func GetJSONName() ([]string, error) {
	files, err := ioutil.ReadDir(config.FolderBathURL)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name()[:len(file.Name())-5])
	}
	return names, nil
}
