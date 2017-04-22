package fileTools

import (
	"os"
	"io/ioutil"
)

/*
文件相关操作封装
@author: yubang
创建于2017年04月22日
 */

/*
读取整个文件内容
@param filePath: 文件路径
@return 文件内容[]byte
 */
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

/*
判断文件是否存在
@param filePath: 文件路径
@return bool
 */
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