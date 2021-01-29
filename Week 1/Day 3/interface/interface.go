package main

import "fmt"

type Hasname interface {
	GetName()
	GetAge()
}

type Person struct {
	name string
	age int
}



func SayHello(person Hasname) {
	fmt.Println("hello", person.GetName(),"anda berumur")
}



func (person Person) GetName() string{
	return person.name
}

func (person Person) GetAge() int{
	return person.age
}

func ups(i int) interface{}{ //apabila method sama sekali tidak dibuat maka interface bisa mereturn apapun tipe datanya
	if i == 1 {
		return 1
	}else if i == 2{
		return true
	}else {
		return ""
	}
}


func main() {
	var person Person 
	person.name = "tika"
	person.age = 40
	
	
	
	SayHello(person)

	var data interface{} = Ups(3)
	fmt.Println(data)


	//interface assertions
	//tanpa switch

	var result interface{} = random()
	var resultString string = result.(int)
	fmt.Println(resultInt)




	switch value := result.(type){
	case string :
		fmt.Println("Program sudah benar bertipe string nilai")
	}
}

