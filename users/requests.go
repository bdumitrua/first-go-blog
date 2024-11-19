package users

import (
	"errors"
	"first-blog-api/utils"
	"net/http"
)

type UpdateUserRequest struct {
	utils.Request
}

func MakeUpdateRequest(w http.ResponseWriter, r *http.Request) *UpdateUserRequest {
	return &UpdateUserRequest{*utils.MakeRequest(w, r)}
}

func (req *UpdateUserRequest) Validate() (map[string]interface{}, error) {
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

	if len([]rune(name)) < 25 {
		http.Error(req.Writer(), "Field 'name' cannot be longer than 25 symbols", http.StatusBadRequest)
		return json, errors.New("field 'name' cannot be longer than 25 symbols")
	}

	return json, nil
}

func (req *UpdateUserRequest) ToDTO() (*UserUpdateDTO, error) {
	reqData, err := req.Validate()
	if err != nil {
		return nil, err
	}

	userDTO := UserUpdateDTO{
		Name: reqData["name"].(string),
	}

	return &userDTO, nil
}
