package main
import "fmt"
func main(){
	var num int
	fmt.Println("enter number ")
	fmt.Scanln(&num)
	for a:=0 ;a<num ;a++{
		fmt.Println(a)
	}
}