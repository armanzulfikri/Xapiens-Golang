package main 

import (
	"fmt"
)

func main(){
    var name  = "Arman"
    var e byte = name[0]
	var eString string =  string(e)
	
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
