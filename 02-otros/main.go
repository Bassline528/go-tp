package notmain

import "fmt"

func main() {
	fmt.Println("Esto es una cadena de texto!")
	fmt.Println("Operacion = ", 2+4)
	fmt.Println(false)
}

func main() {
	const x int = 25000
	x =
		fmt.Println(x)
}

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
}

func main() {
	i := 1
	for i <= 100 {
		fmt.Println(i)
		i = i + 1
	}
}

func main() {
	var i int = 41
	if i%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}
}

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func f(x int) {
	fmt.Println(x)
}

func main() {
	x := 5
	f(x)
}
