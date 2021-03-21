package main

import (
    "fmt"
   "time"
)
func main() {
	
	i:=2
    fmt.Print("write ",i," as : ")

    switch i {
    case 1: 
    	fmt.Print("one\n")
    case 2: 
    	fmt.Print("two\n")
    case 3: 
    	fmt.Print("three\n")
    }
     switch time.Now().Weekday() {
     case time.Saturday, time.Sunday:
     	fmt.Println("it is weekend")
     default :
     	fmt.Println("it is a weekday")
     }

    
}
func hello_world(){

	fmt.Println("hello 2")
}
