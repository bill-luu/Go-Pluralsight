package main

import (
	"net/http"

	"github.com/bill-luu/go-Pluralsight/Go-Getting-Started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
