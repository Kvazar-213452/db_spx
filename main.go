package main

import (
	"fmt"
	func_db "head/func_com"
	"net/http"
)

func main() {
	http.HandleFunc("/check", func_db.Check_db)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Сервер запущено на http://localhost:8080")
}
