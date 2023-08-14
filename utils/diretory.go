package utils

import "os"

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, nil
	}
	// 如果err的错误类型表示文件或目录不存在，os.IsNotExist函数将返回true
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
