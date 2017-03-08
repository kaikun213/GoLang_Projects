
import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var n int
	n0 := 0
	n1 := 1
	depth := 0
	return func() int {
		if (depth == 0){
			n = n0
		} else if (depth == 1) {
			n = n1
		} else {
			n = n0 + n1
			n0 = n1
			n1 = n
		}
		depth++
		return 	n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
