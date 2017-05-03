package docker

/*
调用shell操作docker
@author: yubang
创建于2017年5月3日
 */

import (
	"../../ctsFrame/typeConversionTools"
	"../../ctsFrame/shellTools"
	"fmt"
	"strings"
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

func GetNodeList()[]map[string]string{
	t:= shellTools.RunCommand("docker node ls|awk 'NR!=1{print}'")
	if t == nil{
		return []map[string]string{}
	}
	lines := strings.Split(string(t), "\n")
	result := make([]map[string]string, len(lines))
	for index, line := range lines{
		arrs := strings.Split(line, " ")
		if len(arrs) <5{
			continue
		}
		result[index] = make(map[string]string)
		result[index]["name"] = arrs[3]
		result[index]["status"] = arrs[8]
		result[index]["addr"] = "未知"
	}

	length := 0
	for _, obj := range result{
		if obj != nil{
			length++
		}
	}

	d := make([]map[string]string, length)
	index := 0
	for _, obj := range result{
		if obj != nil {
			d[index] = obj
			index++
		}
	}

	return d
}
