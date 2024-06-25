package integers

import "fmt"

// TODO cleanup this
const englishHelloPrefix = "Hello, "

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

func main() {
	for i := range 10 {
		fmt.Println(Add(i, i))
	}
}
