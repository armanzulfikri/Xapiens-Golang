package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Simple app Calculator!")
	cmd := readLine("Enter Command : [+]add, [-]subtract, [*]multiply, [/]divide")
	fmt.Println(cmd)
	if cmd == "+"{
		num1, num2 := getUserNumbers()
		result := num1 +num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "-" {
		num1, num2 := getUserNumbers()
		result := num1 -num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
		
	} else if cmd == "*" {
		num1, num2 := getUserNumbers()
		result := num1 *num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
		
	} else if cmd == "/" {
		num1, num2 := getUserNumbers()
		result := float32(num1)/float32(num2)
		sResult := fmt.Sprintf("%f", result)
		fmt.Println(sResult)
		
	} else {
		fmt.Println("Invalid input")
	}
}



func readLine(message string) string{
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers()(int, int){
	num1String := readLine("First Number: ")
	num1, err := strconv.Atoi(num1String)
	if err != nil {
		fmt.Println(err)
	}
	num2String := readLine("Second Number:")
	num2, err := strconv.Atoi(num2String)
	if err != nil {
		fmt.Println(err)
	}
	return num1, num2
}