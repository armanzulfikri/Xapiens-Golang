package main 

import (
	"fmt"
)

func main(){
	var i int = 1 
	for i<=20 {
		if i%2 == 0 {
			fmt.Println("Genap")
		}else if i%2 != 0 {
			fmt.Println("Ganjil")
		}else {
			fmt.Println("Somtething wrong")
		}
		i = i + 1;
	}

}
