package main

import (
	"fmt"
	"net/http"

	"my-app/src/routes"
)

func main() {
	r := routes.Routes()

	fmt.Println("Server running on port 5001")
	http.ListenAndServe(":5001", r)
}
