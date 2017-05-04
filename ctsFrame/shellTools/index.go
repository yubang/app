package shellTools

import (
	"os/exec"
	"fmt"
)

/*
shell模块封装
@author: yubang
创建于2017年5月4日
 */

func RunCommand(command string)[]byte{
	f, err := exec.Command("bash", "-c", command).Output()
	fmt.Print(command + "\n")
	if err != nil{
		fmt.Print(err)
		return nil
	}
	return f
}