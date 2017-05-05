package docker

/*
调用shell操作docker
@author: yubang
创建于2017年5月3日
 */

import (
	"../../ctsFrame/typeConversionTools"
	"../../ctsFrame/shellTools"
	"../../ctsFrame/jsonTools"
	"strings"
)

// 创建服务
func CreateService(appId string, nums int, port int, imageName string)bool{
	s := "docker service create --replicas " + typeConversionTools.IntToString(nums) + " --name "+appId+"  -p "+typeConversionTools.IntToString(port)+":80  "+imageName+" /bin/bash /var/start.sh"
	return shellTools.RunCommand(s) != nil
}

// 删除服务
func DeleteService(appId string)bool{
	s := "docker service rm " + appId
	return shellTools.RunCommand(s) != nil
}

// 更新服务镜像
func UpdateImage(appId string, imageName string)bool{
	s := "docker service update --image "+ imageName +" " + appId
	return shellTools.RunCommand(s) != nil
}

// 修改容器信息
func UpdateContainer(appId string, nums int, cpu int, memory int)bool{
	s := "docker service update --replicas " + typeConversionTools.IntToString(nums) + " --limit-memory "+ typeConversionTools.IntToString(memory) +"M --limit-cpu " + typeConversionTools.IntToString(cpu) + " " + appId
	if shellTools.RunCommand(s) == nil{
		return false
	}
	return true
}

// 获取节点信息
func GetNodeInfo(nodeId string)map[string]interface{}{
	t := shellTools.RunCommand("docker node inspect " + nodeId)
	objs := jsonTools.JsonToInterfaceList(t)
	if len(objs) == 1{
		return objs[0].(map[string]interface{})
	}
	return nil
}

// 获取节点列表
func GetNodeList()[]map[string]interface{}{
	t:= shellTools.RunCommand("docker node ls|awk 'NR!=1{print}'")
	if t == nil{
		return []map[string]interface{}{}
	}
	lines := strings.Split(string(t), "\n")
	result := make([]map[string]interface{}, len(lines))
	for index, line := range lines{
		arrs := strings.Split(line, " ")
		if len(arrs) <5{
			continue
		}
		result[index] = GetNodeInfo(arrs[0])
	}

	length := 0
	for _, obj := range result{
		if obj != nil{
			length++
		}
	}

	d := make([]map[string]interface{}, length)
	index := 0
	for _, obj := range result{
		if obj != nil {
			d[index] = obj
			index++
		}
	}

	return d
}

// 获取加入集群命令
func GetJoinCommand()string{
	d := shellTools.RunCommand("docker swarm join-token manager")
	if d == nil{
		return ""
	}
	return string(d)
}

// 移除节点服务器
func DeleteNode(nodeId string)bool{
	return shellTools.RunCommand("docker service rm " + nodeId) != nil
}