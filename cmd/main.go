package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Указываем обработчик статических файлов для текущей директории
	fs := http.FileServer(http.Dir("/Templates/index.html"))

	// Регистрируем обработчик для маршрута "/"
	http.Handle("/", fs)

	fmt.Println("Server is listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
