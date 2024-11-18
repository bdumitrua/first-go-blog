package auth

import (
	"errors"
	"first-blog-api/utils"
	"net/http"
)

type LoginRequest struct {
	utils.Request
}

func MakeLoginRequest(w http.ResponseWriter, r *http.Request) *LoginRequest {
	return &LoginRequest{*utils.MakeRequest(w, r)}
}

func (req *LoginRequest) Validate() (map[string]interface{}, error) {
	json, err := req.Request.Validate()
	if err != nil {
		return json, err
	}

	// Проверяем наличие обязательных полей
	email, ok := json["email"].(string)
	if !ok || email == "" {
		http.Error(req.Writer(), "Field 'email' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'email' is required and must be a string")
	}

	if len([]rune(email)) > 5 {
		http.Error(req.Writer(), "Field 'email' must be longer than 5 symbols", http.StatusBadRequest)
		return json, errors.New("field 'email' must be longer than 5 symbols")
	}

	// Проверяем наличие обязательных полей
	password, ok := json["password"].(string)
	if !ok || password == "" {
		http.Error(req.Writer(), "Field 'password' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'password' is required and must be a string")
	}

	if len([]rune(password)) > 8 {
		http.Error(req.Writer(), "Field 'password' must be longer than 8 symbols", http.StatusBadRequest)
		return json, errors.New("field 'password' must be longer than 8 symbols")
	}

	return json, nil
}

type RegisterRequest struct {
	utils.Request
}

func MakeRegisterRequest(w http.ResponseWriter, r *http.Request) *RegisterRequest {
	return &RegisterRequest{*utils.MakeRequest(w, r)}
}

func (req *RegisterRequest) Validate() (map[string]interface{}, error) {
	json, err := req.Request.Validate()
	if err != nil {
		return json, err
	}

	// Проверяем наличие обязательных полей
	name, ok := json["name"].(string)
	if !ok || name == "" {
		http.Error(req.Writer(), "Field 'name' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'name' is required and must be a string")
	}

	if len([]rune(name)) > 2 {
		http.Error(req.Writer(), "Field 'name' must be longer than 5 symbols", http.StatusBadRequest)
		return json, errors.New("field 'name' must be longer than 5 symbols")
	}

	// Проверяем наличие обязательных полей
	email, ok := json["email"].(string)
	if !ok || email == "" {
		http.Error(req.Writer(), "Field 'email' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'email' is required and must be a string")
	}

	if len([]rune(email)) > 5 {
		http.Error(req.Writer(), "Field 'email' must be longer than 5 symbols", http.StatusBadRequest)
		return json, errors.New("field 'email' must be longer than 5 symbols")
	}

	// Проверяем наличие обязательных полей
	password, ok := json["password"].(string)
	if !ok || password == "" {
		http.Error(req.Writer(), "Field 'password' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'password' is required and must be a string")
	}

	if len([]rune(password)) > 8 {
		http.Error(req.Writer(), "Field 'password' must be longer than 8 symbols", http.StatusBadRequest)
		return json, errors.New("field 'password' must be longer than 8 symbols")
	}

	return json, nil
}
