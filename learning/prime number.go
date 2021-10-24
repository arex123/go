/* run go program : go run nameOfFile.go  */

package main
import "fmt"

func main(){
	fmt.Println("Enter number: ")
	var num int
	fmt.Scanln(&num)

	var count = 0
	for i:=2;i*i<=num;i++{
		if(num%i==0){
			count++;
		}
	}

	if(count==0){
		fmt.Println(num," is prime")
	}else{
		fmt.Println(num ," is not prime number")
	}
}