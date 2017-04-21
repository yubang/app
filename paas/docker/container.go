package docker

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
