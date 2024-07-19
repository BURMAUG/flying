package main

import (
	"flying/internal/model"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	user := model.NewUser("Burmau", *model.NewFlight("Airbus", 12.2))
	fmt.Println(user)
}
