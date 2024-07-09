


package main
import "fmt"

func main() {
	Max()
	fmt.Println(Max)
}

func Max(x, y int) int {
	var max int
	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}