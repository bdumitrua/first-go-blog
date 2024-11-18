package posts

import (
	"errors"
	"first-blog-api/utils"
	"net/http"
)

type CreatePostRequest struct {
	utils.Request
}

func MakeCreateRequest(w http.ResponseWriter, r *http.Request) *CreatePostRequest {
	return &CreatePostRequest{*utils.MakeRequest(w, r)}
}

func (req *CreatePostRequest) Validate() (map[string]interface{}, error) {
	json, err := req.Request.Validate()
	if err != nil {
		return json, err
	}

	// Проверяем наличие обязательных полей
	title, ok := json["title"].(string)
	if !ok || title == "" {
		http.Error(req.Writer(), "Field 'title' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'title' is required and must be a string")
	}

	if len([]rune(title)) > 40 {
		http.Error(req.Writer(), "Field 'title' cannot be longer than 40 symbols", http.StatusBadRequest)
		return json, errors.New("field 'title' cannot be longer than 40 symbols")
	}

	content, ok := json["content"].(string)
	if !ok || content == "" {
		http.Error(req.Writer(), "Field 'content' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'content' is required and must be a string")
	}

	if len([]rune(content)) > 255 {
		http.Error(req.Writer(), "Field 'content' cannot be longer than 255 symbols", http.StatusBadRequest)
		return json, errors.New("field 'content' cannot be longer than 255 symbols")
	}

	return json, nil
}

func (req *CreatePostRequest) ToDTO() (*PostDTO, error) {
	reqData, err := req.Validate()
	if err != nil {
		return nil, err
	}

	postDTO := PostDTO{
		Title:   reqData["title"].(string),
		Content: reqData["content"].(string),
	}

	return &postDTO, nil
}

// Same logic, so...

type UpdatePostRequest struct {
	utils.Request
}

func MakeUpdateRequest(w http.ResponseWriter, r *http.Request) *UpdatePostRequest {
	return &UpdatePostRequest{*utils.MakeRequest(w, r)}
}

func (req *UpdatePostRequest) Validate() (map[string]interface{}, error) {
	json, err := req.Request.Validate()
	if err != nil {
		return json, err
	}

	// Проверяем наличие обязательных полей
	title, ok := json["title"].(string)
	if !ok || title == "" {
		http.Error(req.Writer(), "Field 'title' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'title' is required and must be a string")
	}

	if len([]rune(title)) > 40 {
		http.Error(req.Writer(), "Field 'title' cannot be longer than 40 symbols", http.StatusBadRequest)
		return json, errors.New("field 'title' cannot be longer than 40 symbols")
	}

	content, ok := json["content"].(string)
	if !ok || content == "" {
		http.Error(req.Writer(), "Field 'content' is required and must be a string", http.StatusBadRequest)
		return json, errors.New("field 'content' is required and must be a string")
	}

	if len([]rune(content)) > 255 {
		http.Error(req.Writer(), "Field 'content' cannot be longer than 255 symbols", http.StatusBadRequest)
		return json, errors.New("field 'content' cannot be longer than 255 symbols")
	}

	return json, nil
}

func (req *UpdatePostRequest) ToDTO() (*PostDTO, error) {
	reqData, err := req.Validate()
	if err != nil {
		return nil, err
	}

	postDTO := PostDTO{
		Title:   reqData["title"].(string),
		Content: reqData["content"].(string),
	}

	return &postDTO, nil
}
