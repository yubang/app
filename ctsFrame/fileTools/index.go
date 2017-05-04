package fileTools

import (
	"os"
	"io/ioutil"
	"../utilTools"
	"errors"
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


func MakeDirs(dirPath string)bool{
	return os.MkdirAll(dirPath, 0777) == nil
}

func MakeTempDir()(string, error){
	// todo: 现在仅支持linux
	dirPath := "/tmp/ctsFrame/tmp/" + utilTools.GetToken32()
	if !MakeDirs(dirPath){
		return "", errors.New("无法创建临时文件夹")
	}
	return dirPath, nil
}

func WriteNewFile(filePath string, fileContent []byte)bool{
	return ioutil.WriteFile(filePath, fileContent, 0777) == nil
}

func RemoveDir(dirPath string)bool{
	return os.RemoveAll(dirPath) == nil
}