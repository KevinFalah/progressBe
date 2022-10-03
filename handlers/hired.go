package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	hireddto "waysgallery/dto/hired"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerHired struct {
	HiredRepository repositories.HiredRepository
}

// Create `path_file` Global variable here ...

func HandlerHired(HiredRepository repositories.HiredRepository) *handlerHired {
	return &handlerHired{HiredRepository}
}

func (h *handlerHired) FindHireds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hireds, err := h.HiredRepository.FindHireds()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: hireds}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHired) GetHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var hired models.Hired
	hired, err := h.HiredRepository.GetHired(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	orderTo, _ := h.HiredRepository.GetOrderTo(hired.OrderToID)

	hired, _ = h.HiredRepository.GetHired(hired.ID)

	orderToResponse := models.HiredResponse{
		ID:             hired.ID,
		Title:          hired.Title,
		DescriptionJob: hired.DescriptionJob,
		StartProject:   hired.StartProject,
		EndProject:     hired.EndProject,
		Price:          hired.Price,
		User:           hired.User,
		OrderTo: models.OrderToHired{
			ID:       orderTo.ID,
			FullName: orderTo.FullName,
			Email:    orderTo.Email,
		},
	}

	// Create Embed Path File on Image property here ...

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: orderToResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHired) CreateHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	request := new(hireddto.HiredRequest)

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

	hired := models.Hired{
		Title:          request.Title,
		DescriptionJob: request.DescriptionJob,
		StartProject:   request.StartProject,
		EndProject:     request.EndProject,
		Price:          request.Price,
		UserID:         userId,
		OrderToID:      request.OrderToID,
	}

	// err := mysql.DB.Create(&hired).Error
	hired, err = h.HiredRepository.CreateHired(hired)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	orderTo, _ := h.HiredRepository.GetOrderTo(hired.OrderToID)

	hired, _ = h.HiredRepository.GetHired(hired.ID)

	orderToResponse := models.HiredResponse{
		ID:             hired.ID,
		Title:          hired.Title,
		DescriptionJob: hired.DescriptionJob,
		StartProject:   hired.StartProject,
		EndProject:     hired.EndProject,
		Price:          hired.Price,
		User:           hired.User,
		OrderTo: models.OrderToHired{
			ID:       orderTo.ID,
			FullName: orderTo.FullName,
			Email:    orderTo.Email,
		},
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: orderToResponse}
	json.NewEncoder(w).Encode(response)
}

func convertResponseHired(u models.Hired) models.HiredResponse {
	return models.HiredResponse{
		ID:             u.ID,
		Title:          u.Title,
		DescriptionJob: u.DescriptionJob,
		StartProject:   u.StartProject,
		EndProject:     u.EndProject,
		Price:          u.Price,
		// UserID:         0,
		User: u.User,
		// OrderToID:      0,
		OrderTo: models.OrderToHired{ID: u.OrderTo.ID, FullName: u.OrderTo.Email, Email: u.OrderTo.Email},
		// CreatedAt:      time.Time{},
		// UpdatedAt:      time.Time{},
	}
}
