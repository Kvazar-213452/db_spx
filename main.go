package main

import (
	"fmt"
	func_db "head/func_com"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]

	http.HandleFunc("/check", func_db.Check_server)
	http.HandleFunc("/check_db", func_db.Check_db)
	http.HandleFunc("/add", func_db.Add_db)
	http.HandleFunc("/del", func_db.Del_data_db)
	http.HandleFunc("/del_db", func_db.Del_all_data_db)

	fmt.Printf("Сервер запущено на http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Помилка при запуску сервера:", err)
	}
}
