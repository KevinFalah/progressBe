package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	followingdto "waysgallery/dto/following"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerFollowing struct {
	FollowingRepository repositories.FollowingRepository
}

// Create `path_file` Global variable here ...

func HandlerFollowing(FollowingRepository repositories.FollowingRepository) *handlerFollowing {
	return &handlerFollowing{FollowingRepository}
}

func (h *handlerFollowing) FindFollowings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	followings, err := h.FollowingRepository.FindFollowings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: followings}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFollowing) GetFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var following models.Following
	following, err := h.FollowingRepository.GetFollowing(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: following}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFollowing) CreateFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	request := new(followingdto.FollowingRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	following := models.Following{
		UserID:      userId,
		FollowingID: request.FollowingID,
	}

	// err := mysql.DB.Create(&Following).Error
	following, err = h.FollowingRepository.CreateFollowing(following)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: following}
	json.NewEncoder(w).Encode(response)
}

func convertResponseFollowing(u models.Following) models.FollowingResponse {
	return models.FollowingResponse{
		ID:          u.ID,
		UserID:      0,
		User:        u.User,
		FollowingID: 0,
		Following:   models.OrderToHired{ID: u.Following.ID, FullName: u.Following.Email, Email: u.Following.Email},
	}
}
