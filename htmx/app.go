package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "<p>GET request received</p>")
		case "POST":
			fmt.Fprintf(w, "<p>POST request received</p>")
		case "PUT":
			fmt.Fprintf(w, "<p>PUT request received</p>")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
