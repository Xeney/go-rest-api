package handlers

import (
	"encoding/json"
	"http-server/models"
	"log"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading page /update")

	if r.Method != http.MethodPut {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	request := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Invalid JSON: ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if _, exists := models.UsersBase[request.Name]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	models.UsersBase[request.Name] = *request

	w.Header().Set("Content-Type", "application/json")
	models.SendResponse(w, "success", "User "+request.Name+" updated", http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading page /delete")

	if r.Method != http.MethodDelete {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	request := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Invalid JSON: ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if _, exists := models.UsersBase[request.Name]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(models.UsersBase, request.Name)

	w.Header().Set("Content-Type", "application/json")
	models.SendResponse(w, "success", "User "+request.Name+" deleted", http.StatusOK)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading page /create")

	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	request := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Invalid JSON: ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if request.Name == "" && request.Age > 0 {
		log.Println("Bad Request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, exists := models.UsersBase[request.Name]; exists {
		// Лучше вообще никакую ошибку не сообщать, ибо это потенциальная
		// дыра в безопасности.
		// Злоумышленник таким образом может через тебя прогнать свою базу
		// адресов и узнать, какие люди у тебя зарегистрированы,
		// что может быть нежелательно
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	models.UsersBase[request.Name] = *request

	w.Header().Set("Content-Type", "application/json")
	models.SendResponse(w, "success", "User "+request.Name+" created", http.StatusOK)
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("loading page /hello")
	w.Header().Set("Content-Type", "application/json")

	models.SendResponse(w, "success", "Hello, World!", http.StatusOK)
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("loading page /health")

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Println("Error write_byte: ", err)
	}
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loading page /greet")
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}
	models.SendResponse(w, "success", "Hello, "+name+"!", http.StatusOK)
}
