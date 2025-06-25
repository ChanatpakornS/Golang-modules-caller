package main

import (
	"fmt"

	"github.com/ChanatpakornS/Golang-private-mods-demo/domain/hello"
	"github.com/ChanatpakornS/Golang-private-mods-demo/ports"
)

func main() {

	repo := hello.NewTestRepo()
	handler := ports.NewTestHandeler(repo)
	result := handler.Handle()

	fmt.Println("Demo application started!", result)
}
