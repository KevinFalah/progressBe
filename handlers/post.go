package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	postdto "waysgallery/dto/post"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

// Create `path_file` Global variable here ...

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) FindPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := h.PostRepository.FindPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	// filePath := os.Getenv("PATH_FILE")

	// for i, p := range posts {
	// 	posts[i].Photo = os.Getenv("PATH_FILE") + p.Photo
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: posts}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var post models.Post
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	// literatur.Photo = os.Getenv("PATH_FILE") + literatur.Photo

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var post models.Post
// 	post, err := h.PostRepository.GetPost(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	filePath := os.Getenv("PATH_FILE")

// 	// Create Embed Path File on Image property here ...
// 	postResponse := postdto.PostResponse{
// 		ID:          post.ID,
// 		Photo:       filePath + post.Photo,
// 		Title:       post.Title,
// 		Description: post.Description,
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: postResponse}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	// Get dataFile from midleware and store to filename variable here ...
	request := postdto.PostRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post := models.Post{
		Title:       request.Title,
		Description: request.Description,
		Photo:       os.Getenv("PATH_FILE") + filename,
		UserID:      userId,
	}

	// err := mysql.DB.Create(&post).Error
	post, err = h.PostRepository.CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post, _ = h.PostRepository.GetPost(post.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func convertResponsePost(u models.Post) models.PostResponse {
	return models.PostResponse{
		ID:          0,
		Title:       u.Title,
		Description: u.Description,
		Photo:       u.Photo,
		User:        u.User,
	}
}

// package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	postdto "waysgallery/dto/post"
// 	dto "waysgallery/dto/result"
// 	"waysgallery/models"
// 	"waysgallery/repositories"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/gorilla/mux"
// )

// type handlerPost struct {
// 	PostRepository repositories.PostRepository
// }

// var path_file = "PATH_FILE"

// func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
// 	return &handlerPost{PostRepository}
// }

// func (h *handlerPost) FindPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	posts, err := h.PostRepository.FindPosts()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Untuk mengembed path file di property photo
// 	for i, p := range posts {
// 		posts[i].Photo = os.Getenv("PATH_FILE") + p.Photo
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: posts}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var post models.Post
// 	post, err := h.PostRepository.GetPost(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// path untuk membuat api file image
// 	post.Photo = os.Getenv("PATH_FILE") + post.Photo

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponsePost(post)}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	userId := int(userInfo["id"].(float64))

// 	// Variable untuk memanggil uploadFile
// 	dataContex := r.Context().Value("image")
// 	filename := dataContex.(string)

// 	// category_id, _ := strconv.Atoi(r.FormValue("category_id"))
// 	request := postdto.CreatePostRequest{
// 		Title:       r.FormValue("title"),
// 		Description: r.FormValue("description"),
// 		Photo:       filename,
// 		UserID:      userId,
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	post := models.Post{
// 		Title:       request.Title,
// 		Description: request.Description,
// 		Photo:       filename,
// 		UserID:      request.UserID,
// 	}

// 	// err := mysql.DB.Create(&Post).Error
// 	post, err = h.PostRepository.CreatePost(post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// post, _ = h.PostRepository.GetPost(post.ID)

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlerPost) UpdatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(postdto.UpdatePostRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	post, err := h.PostRepository.GetPost(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	if request.Title != "" {
// 		post.Title = request.Title
// 	}

// 	if request.Description != "" {
// 		post.Description = request.Description
// 	}

// 	if request.Description != "" {
// 		post.Description = request.Description

// 		data, err := h.PostRepository.UpdatePost(post)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponsePost(data)}
// 		json.NewEncoder(w).Encode(response)
// 	}
// }

// func (h *handlerPost) DeletePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	post, err := h.PostRepository.GetPost(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	data, err := h.PostRepository.DeletePost(post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponsePost(data)}
// 	json.NewEncoder(w).Encode(response)
// }

// func convertResponsePost(u models.Post) models.Post {
// 	return models.Post{
// 		ID:          u.ID,
// 		Title:       u.Title,
// 		Description: u.Description,
// 		Photo:       u.Photo,
// 	}
// }
