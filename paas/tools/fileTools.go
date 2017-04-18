package tools

import (
	"os"
	"io/ioutil"
)

/*
读取文件的工具类
@author: yubang
 */


// 读取整个文件内容
func ReadFromFile(filePath string) []byte{

	if !CheckFileExist(filePath){
		return nil
	}

	fi,err := os.Open(filePath)
	defer fi.Close()

	if err != nil{
		return nil
	}

	fd, err := ioutil.ReadAll(fi)
	if err != nil{
		return nil
	}
	return fd
}

// 判断文件是否存在
func CheckFileExist(filePath string) bool{
	_, err := os.Stat(filePath)
	if err != nil{
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}