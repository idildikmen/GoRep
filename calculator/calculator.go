package main

import "fmt"


func main(){
	num1 := 0
	num2 := 0
	calculate := ""

	fmt.Println("Enter The First Number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter The Second Number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter (1) for summation, (2) for substraction, (3) for multiplication and (4) for division: ")
	fmt.Scanln(&calculate)

	if calculate == "1"{
		fmt.Println(num1+num2)
	}else if calculate == "2"{
		fmt.Println(num1-num2)
	}else if calculate == "3"{
		fmt.Println(num1*num2)
	}else {
		fmt.Println(num1/num2)
	}

}


