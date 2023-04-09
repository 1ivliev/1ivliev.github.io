package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port string

	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	http.HandleFunc("/get-info", getHandler)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

type Response struct {
	Status  bool `json:"status"`
	Result Result `json:"result"`
}

type BrokenResponse struct {
	Status bool `json:"status"`
	Error Error `json:"error"`
}

type Result struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Role string `json:"role"`
	Email string `json:"email"`
	Resume string `json:"resume_link"`
}

type Error struct {
	Error string `json:"error_message"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method != "GET":
		response := BrokenResponse{false, Error{"Method not allowed"}}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
	case r.Method == "GET":
		response := Response{true, Result{"Kirill Ivliev", 26, "Automation QA Engineer", "1ivliev@mail.ru", "https://resume.io/r/QiP44dXa3",}}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
