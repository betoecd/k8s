package main

import "net/http"

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("10 bilhões por cento de certeza!!!"))
}
