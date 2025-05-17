package main

import "fmt"

func main() {
	var listafrutas = [4]string{"Pera", "Naranja"}
	fmt.Println(listafrutas[1])

	listapaises := []string{"Peru", "Chile", "Venezuela"}
	fmt.Println(listapaises)

	listapaises = append(listapaises, "Brasil")
	fmt.Println(listapaises)

	listapaises2 := listapaises[1:3]
	fmt.Println(listapaises2)

	listapaises3 := listapaises[2:]
	fmt.Println(listapaises3)

	listapaises4 := listapaises[:3]
	fmt.Println(listapaises4)

}
