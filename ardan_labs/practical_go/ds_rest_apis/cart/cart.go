package main

import "fmt"

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println("len:", len(cart))
	fmt.Println("cart[1]:", cart[1])

	for i, c := range cart {
		fmt.Println(i, c)
	}

	cart = append(cart, "milk")
	fmt.Println(cart)

	fruit := cart[:3]
	fmt.Println("fruit:", fruit)
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit)
	fmt.Println("fruit:", cart)
}
