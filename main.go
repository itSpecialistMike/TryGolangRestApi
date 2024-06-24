package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура ответа
type Response struct {
	Message string `json:"message"`
}

// Обработчик запроса на "/hello_world"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем ответ
	response := Response{"Hello, world!"}

	// Преобразуем ответ в JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправляем ответ
	w.Write(jsonResponse)
}

func main() {
	// Регистрируем обработчик для "/hello_world"
	http.HandleFunc("/hello_world", helloHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080/hello_world")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
