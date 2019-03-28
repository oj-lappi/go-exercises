package main

import(
	"fmt"
)

func main(){
	rekursera("")
}

func rekursera(indentering string){
	if len(indentering) == 10{
		return
	}
	fmt.Println(indentering+"dyker in...",len(indentering))
	rekursera(indentering+" ")
	fmt.Println(indentering+"dyker ut...",len(indentering))
}
