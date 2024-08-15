package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Message string
		}{
			Message: "Hello World!",
		}

		tmpl.Execute(w, data)

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

	fmt.Println("Started http server on port :8080")
	http.ListenAndServe(":8080", nil)
}
