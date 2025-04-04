package main

import "fmt"

func main() {
	var suma int = 0

	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			suma = suma + 1
		}
	}

	fmt.Println("La suma de los primeros 100 numero es ", suma)

}
