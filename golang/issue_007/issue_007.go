package issue007

import "fmt"

func doubleAndPrint(fn func(int) int) func(int) int {
	return func(x int) int {
		result := fn(x)
		fmt.Println("Result:", result*2)
		return result
	}
}

func square(x int) int {
	return x * x
}

func Facde() {
	doubleSquare := doubleAndPrint(square)
	res := doubleSquare(3) // 18 (3*3*2)
	fmt.Println(res)
}
