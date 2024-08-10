func main(){
	http.HandleFunc("/", func(w http.Response))

	if r.Method == "GET" {
		fmt.Fprintf(w, "<p>GET</p>")	
	}

	if r.Method == "POST" {
		fmt.Fprintf(w, "<p>POST</p>")	
	}
}