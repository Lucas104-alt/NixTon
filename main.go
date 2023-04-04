package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	nixton := "NixTon 1.0.0 "
	fmt.Print(nixton)
	fmt.Print("Creado Por Lucas11 ")
	fmt.Println(" UUID:", uuid.New().String())
}
