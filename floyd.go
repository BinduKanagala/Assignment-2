package main
 
import "fmt"
 
func main() {

	var n int
	var a int = 1
	fmt.Print("Enter the no of rows of floyd's triangle to print: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		for c := 1; c <= i; c++ {
		
			fmt.Printf(" %d",a)
			a++
		}
		fmt.Println("")
	}
}