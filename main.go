package main
import "fmt"
func main() {
	fmt.Println("FizzBuzz Program - feature1")
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Print(i, " ")
		}
	}
}