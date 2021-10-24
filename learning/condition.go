package main
import "fmt"
func main(){
	var a int 
	fmt.Println("enter number")
	fmt.Scanln(&a)
	if a%2==0 {
		fmt.Println(a)
	}else {
		fmt.Println(a+1)
	}
	
}