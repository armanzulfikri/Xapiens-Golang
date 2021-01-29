package main

import "fmt"

//basic fungsi
func sayHello(){
	fmt.Println("Hello Xapiens")
}

func sayName(name string)  {
	fmt.Println("Hello my name :", name)
	
}


//fungsi dengan return 
func sayReturn(name string) string  {
	return "hello :" + name
 
	
}

//fungsi dengan multiple return
func sayMultipleReturn() (string,string,string){
	return "ngalik","sleman","DIY"
} 



//fungsi dengan multiple alias
func sayMultipleReturnAlias() (kec,kab,prov string){
	kec = "ngalik"
	kab = "sleman"
	prov = "yogyakarta"

	return 
} 

//fungsi variadic
// func sumAll(pengali int,numbers ...int) int{
// 	total := 0

// 	for_, value := range numbers{
// 			total += value
// 	}

// 	return total * pengali
	
// }


//fungsi sebagai parameter
type Filter func (string) string 


func textFilter(name string, filter Filter){
	filtered := filter(name)
	fmt.Println("Hello", filtered)
}

func filter (text string) string{
	if text == "go-blok"{
		return "......"
	}else {
		return text
	}
}

//anonymous function
type Blacklist func(string) bool 
func registerUser(name string, blacklist Blacklist){
	if(blacklist(name)){
		fmt.Println("are you blocked", name)
	}else{
		fmt.Println("welcome",name)
	}
}

//recrusive
func factorial(value int)int{
	if value == 1{
		return 1
	}else {
		return value * factorial(value-1)
	}
}

func main() {
	fmt.Println(5*4*3*2*1)
	fmt.Println(factorial(5))





var getMinMax = func(numbers ...int)(int,int){
	var min,max int
	for _, v := range numbers{
		if min > v{
			min = v
		}
		if max < v{
			max = v
		}
	}
	return min, max
}

min,max := getMinMax(1,2,3,4,5,6,7,8,9)
fmt.Println(min,max)

	blacklist := func(name string)bool{
		return name == "admin"
	} 

	registerUser("admin", blacklist)
	registerUser("arman", func (name string) bool  {
		return name == "admin"
	})



	sayHello()

	sayName("Arman")

	fmt.Println(sayReturn("arman"))
	fmt.Println(sayMultipleReturn())
	kec,kab,prov := sayMultipleReturn()
	fmt.Println(kec,kab,prov)

	fmt.Println(sayMultipleReturnAlias())
	// fmt.Println(sumAll(2,3,4,5,6,6))



}