package handlers

import (
	"bytes"
	"http-server/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест для HelloHandler
func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"success","message":"Hello, World!"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Тест для HealthHandler
func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Тест для GreetHandler
func TestGreetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/greet?name=Alice", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GreetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"success","message":"Hello, Alice!"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Тест для CreateHandler
func TestCreateHandler(t *testing.T) {
	requestBody := `{"name": "Alice", "age": 25}`
	req, err := http.NewRequest("POST", "/create", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"success","message":"User Alice created"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if _, exists := models.UsersBase["Alice"]; !exists {
		t.Errorf("user Alice was not added to the database")
	}
}

// Тест для UpdateHandler
func TestUpdateHandler(t *testing.T) {
	// Добавляем пользователя для обновления
	models.UsersBase["Bob"] = models.User{Name: "Bob", Age: 30}

	requestBody := `{"name": "Bob", "age": 35}`
	req, err := http.NewRequest("PUT", "/update", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"success","message":"User Bob updated"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if models.UsersBase["Bob"].Age != 35 {
		t.Errorf("user Bob's age was not updated: got %v want %v",
			models.UsersBase["Bob"].Age, 35)
	}
}

// Тест для DeleteHandler
func TestDeleteHandler(t *testing.T) {
	// Добавляем пользователя для удаления
	models.UsersBase["Charlie"] = models.User{Name: "Charlie", Age: 40}

	requestBody := `{"name": "Charlie"}`
	req, err := http.NewRequest("DELETE", "/delete", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"success","message":"User Charlie deleted"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if _, exists := models.UsersBase["Charlie"]; exists {
		t.Errorf("user Charlie was not deleted from the database")
	}
}
