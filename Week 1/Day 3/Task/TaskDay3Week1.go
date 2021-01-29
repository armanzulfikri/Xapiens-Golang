package main  
  
import "fmt"  

func sum(number1 int, number2 int) int{
	total := number1 + number2
	return total
}

func pembagian(number1 int, number2 int) int{
	bagi := number1/number2
	return bagi
}

func perkalian(number1 int, number2 int) int{
	kali := number1 * number2
	return kali
}

func pengurangan(number1 int, number2 int) int{
	kurang := number1 - number2
	return kurang
}


func main() {  
 var operator,name string  
 var number1, number2 int  
 fmt.Println("Please enter name :")
 fmt.Scanln(&name)
 fmt.Print("Please enter First number: ")  
 fmt.Scanln(&number1)  
 fmt.Print("Please enter Second number: ")  
 fmt.Scanln(&number2)  
 fmt.Print("Please enter Operator (+,-,/,%,*):")  
 fmt.Scanln(&operator)  
 output := 0  
 switch operator {  
 case "+":  
  sum(number1,number2)
  break  
 case "-":  
  output = number1 - number2  
 case "*":  
  output = number1 * number2  
 case "/":  
  output = number1 / number2  
 case "%":  
  output = number1 % number2  
 default:  
  fmt.Println("Invalid Operation")   
 }  
 fmt.Printf("%d %s %d = %d", number1, operator, number2, output)  
}  




