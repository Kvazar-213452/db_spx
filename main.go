package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не дозволено", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Сервер запущено на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Помилка при запуску сервера:", err)
	}
}
