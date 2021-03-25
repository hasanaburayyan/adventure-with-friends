package main

import (
	"adventure-with-friends/src/models"
	"fmt"
)

func main() {
	fmt.Println("Starting Adventure With Friends Server...")
	// The Magic
	fmt.Println("If I were a real application I would continue running and serving requests!")
	fmt.Println("Instead Ill just make a Player and show you his info!")
	p := models.Player{
		Name: "Cool",
		Level: 1,
	}
	fmt.Println(p);
	fmt.Println("Stopping Adventure With Friends Server...")
}
