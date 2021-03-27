package main

import (
	"fmt"

	"github.com/bill-luu/go-Pluralsight/Go-Getting-Started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Bill",
		LastName:  "Luu",
	}
	fmt.Println(u)
}
