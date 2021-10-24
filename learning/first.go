//package main

//import "fmt"

//func main() {
//    fmt.Println("Hello, World!")
//}


package main

import "fmt"
func main(){

	fmt.Println("Enter string")

	var first string

	fmt.Scanln(&first)

	fmt.Println("Enter final string")

	var second string

	fmt.Scanln(&second)


	var com string = first + second

	fmt.Println(com)







}