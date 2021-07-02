package resource_utils

import "os"

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
