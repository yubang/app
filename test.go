package main

import (
	"./ctsFrame/stringTools"
	"fmt"
)

func main(){
	v := stringTools.SubString("1核", 0, (len("1核") /2)-1 )
	fmt.Print(len("1核"))
	fmt.Print(v)
}