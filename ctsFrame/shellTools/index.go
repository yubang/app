package shellTools

import (
	"os/exec"
	"fmt"
	"time"
	"errors"
	"../typeConversionTools"
)

/*
shell模块封装
@author: yubang
创建于2017年5月4日
 */

func KillAllProcess(pid int){
	command := "pstree "+typeConversionTools.IntToString(pid)+" -p | awk -F\"[()]\" '{for(i=0;i<=NF;i++)if($i~/[0-9]+/)print $i}'|awk 'NR!=1{print$1}'|xargs kill"
	fmt.Print(command)
	exec.Command(command).Output()
}

func executeSheel(cmd *exec.Cmd, shellChan chan bool, b chan []byte, e chan error)([]byte, error){
	f, err := cmd.Output()
	if err != nil{
		fmt.Print(err)
		shellChan <- true
		b <- nil
		e <- err
		return nil, err
	}
	shellChan <- true
	b <- f
	e <- nil
	return f, nil
}



func RunCommand(command string, timeout time.Duration)([]byte, error){
	fmt.Print(command + "\n")
	cmd := exec.Command("bash", "-c", command)

	// 启动shell脚本
	shellChan := make(chan bool)
	b := make(chan []byte)
	e := make(chan error)

	go executeSheel(cmd, shellChan, b, e)

	select {

		case <- shellChan:
			// 完成shell执行
			return <-b, <-e

		case <- time.After(timeout):
			KillAllProcess(cmd.Process.Pid)
			cmd.Process.Kill()
			return nil, errors.New("执行shell命令超时")



	}

}