package jutils

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitee.com/hhvy/golib/jerrors"
)

func FileIsExists(filepath string) bool {
	_, err := os.Lstat(filepath)
	return !os.IsNotExist(err)
}

func ReadFileAll(path string) (string, error) {

	if !FileIsExists(path) {
		return "", jerrors.NewError("文件不存在")
	}

	f, err := os.Open(path)
	if err != nil {
		s := fmt.Sprintf("读取文件失败:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		s := fmt.Sprintf("获取文件内容失败:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	return string(fd), nil
}
