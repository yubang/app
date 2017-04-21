package docker

import (
	"net"
	"strconv"
)

/*
docker的容器操作相关封装
@author: yubang
 */

// 获取一种镜像的容器
func GetATypeOfImageContainerNumber(imageName string)int{
	return 0
}

// 启动一个容器
func StartAContainer(imageName string, port int, dockerPort int) bool{
	return true
}

// 删除一个容器
func RemoveAContainer(containerId string) bool{
	return true
}

// 获取所有容器信息
func GetAllContainerList() []map[string]interface{}{ // return [{"containerId": 容器id, "imageName": 镜像名字, "port": 使用的端口}]
	return nil
}

// 获取一个容器
func GetAContainer(imageName string) string{
	return ""
}

// 检测端口是否可用
func CheckPort(port int)bool{
	_, err := net.Dial("tcp", "127.0.0.1:" + strconv.Itoa(port))
	if err != nil{
		return true
	}
	return false
}

// 获取一个可用端口
func GetAbleUserPort()int{
	for port := 20000;port <= 65535;port++{
		if CheckPort(port){
			return port
		}
	}
	return 0
}
