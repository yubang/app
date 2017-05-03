package docker

/*
调用shell操作docker
@author: yubang
创建于2017年5月3日
 */

import (
	"../../ctsFrame/typeConversionTools"
	"fmt"
)


func CreateService(appId string, nums int, port int, imageName string)bool{
	s := "docker service create --replicas " + typeConversionTools.IntToString(nums) + " --name "+appId+"  -p "+typeConversionTools.IntToString(port)+":80  "+imageName+" /var/start.sh"
	fmt.Print(s + "\n")
	return true
}

func UpdateImage(appId string, imageName string)bool{
	s := "docker service update --image "+ imageName +" " + appId
	fmt.Print(s + "\n")
	return true
}

func UpdateContainer(appId string, nums int, cpu int, memory int)bool{
	s := "docker service update  --reserve-memory "+ typeConversionTools.IntToString(memory) +"M --limit-cpu " + typeConversionTools.IntToString(cpu) + " " + appId
	s = "docker service scale " + appId + "=" + typeConversionTools.IntToString(nums)
	fmt.Print(s + "\n")
	return true
}