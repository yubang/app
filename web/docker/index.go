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
	"../../ctsFrame/cacheTools"
	"../../ctsFrame/timeTools"
)


type ShellStruct struct{
	Client cacheTools.RedisClientObject
}

func (obj *ShellStruct)ExecShell(command string)[]byte{
	d, err := shellTools.RunCommand(command)
	if err != nil{
		// 记录日志
		log := jsonTools.InterfaceToJson(map[string]interface{}{
			"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
			"command": command,
			"content": err.Error(),
		})
		obj.Client.GetRedisClient().LPush("paas_error_shell_list", log)
		return nil
	}
	return d
}


type DockerStruct struct {
	Client cacheTools.RedisClientObject
}

func (obj *DockerStruct)getClient()*ShellStruct{
	shell := ShellStruct{obj.Client}
	return &shell
}

// 创建服务
func (obj *DockerStruct)CreateService(appId string, nums int, port int, imageName string)bool{
	s := "docker service create --replicas " + typeConversionTools.IntToString(nums) + " --name "+appId+"  -p "+typeConversionTools.IntToString(port)+":80 --network " + appId + " " +imageName+" /bin/bash /var/start.sh"
	return obj.getClient().ExecShell(s) != nil
}

// 删除服务
func (obj *DockerStruct)DeleteService(appId string)bool{
	s := "docker service rm " + appId
	return obj.getClient().ExecShell(s) != nil
}

// 更新服务镜像
func (obj *DockerStruct)UpdateImage(appId string, imageName string)bool{
	s := "docker service update --image "+ imageName +" " + appId
	return obj.getClient().ExecShell(s) != nil
}

// 修改容器信息
func (obj *DockerStruct)UpdateContainer(appId string, nums int, cpu int, memory int)bool{
	s := "docker service update --replicas " + typeConversionTools.IntToString(nums) + " --limit-memory "+ typeConversionTools.IntToString(memory) +"M --limit-cpu " + typeConversionTools.IntToString(cpu) + " " + appId
	if obj.getClient().ExecShell(s) == nil{
		return false
	}
	return true
}

// 获取节点信息
func (obj *DockerStruct)GetNodeInfo(nodeId string)map[string]interface{}{
	t := obj.getClient().ExecShell("docker node inspect " + nodeId)
	objs := jsonTools.JsonToInterfaceList(t)
	if len(objs) == 1{
		return objs[0].(map[string]interface{})
	}
	return nil
}

// 获取节点列表
func (obj *DockerStruct)GetNodeList()[]map[string]interface{}{
	t:= obj.getClient().ExecShell("docker node ls|awk 'NR!=1{print}'")
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
		result[index] = obj.GetNodeInfo(arrs[0])
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
func (obj *DockerStruct)GetJoinCommand()string{
	d := obj.getClient().ExecShell("docker swarm join-token worker")
	if d == nil{
		return ""
	}
	s := strings.Replace(string(d), "\\", "", -1)
	s = strings.Replace(s, "To add a worker to this swarm, run the following command:", "", -1)

	return s
}

// 移除节点服务器
func (obj *DockerStruct)DeleteNode(nodeId string)bool{
	return obj.getClient().ExecShell("docker node rm " + nodeId) != nil
}

// 创建网络
func (obj *DockerStruct)CreateNet(netName string)bool{
	return obj.getClient().ExecShell("docker network create -d overlay " + netName) != nil
}

// 删除网络
func (obj *DockerStruct)DeleteNet(netName string)bool{
	return obj.getClient().ExecShell("docker network rm " + netName) != nil
}