package main

import (
	"net/http"
)

func main() {

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			/// http.ServeFile(w, r, "./index.html")
			//http.ServeFile(w, r, r.URL.Path[1:])
		})
	*/

	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.ListenAndServe(":8000", nil)
}
