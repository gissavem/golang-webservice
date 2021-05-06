package main

import (
	"fmt"

	"github.com/gissavem/golang-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "John",
		LastName:  "Andersson",
	}
	fmt.Println(u)
}
