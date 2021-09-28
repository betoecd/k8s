package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s. I'm %s years old.", name, age)
}
