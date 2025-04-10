package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func readPage(pageName string) []byte {
	 page, err := os.ReadFile(pageName)
	 if err != nil {
		 log.Fatal(err)
	 }
	return page
}

func homePage(w http.ResponseWriter, r *http.Request) {
	page := readPage("home.html")
	_, err := w.Write(page)
	if err != nil {
		fmt.Println("Ошибка при открытии страницы:" ,err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := Message{
		Status: "success",
		Message: "JSON",
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Ошибка при отправки JSON :", err)
	}
}

func main() {
	http.HandleFunc("/home",homePage)
	http.HandleFunc("/api",api)
	err := http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)
}
