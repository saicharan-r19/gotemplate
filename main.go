package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Hello("sai").Render(r.Context(), w)
	})

	http.ListenAndServe(":3000", nil)
}
