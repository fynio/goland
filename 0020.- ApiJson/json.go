package main

import (
	"encoding/json"
	"net/http"
	// "encoding/json"
)

type curso struct {
	Title        string `json:"titulo"`
	NumeroVideos int
}

type cursos []curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cursoss := cursos{
			curso{"Programacion", 20},
			curso{"Hacking", 20},
		}

		json.NewEncoder(w).Encode(cursoss)
	})

	http.ListenAndServe(":8001", nil)
}
