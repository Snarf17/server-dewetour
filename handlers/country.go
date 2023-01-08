package handlers

import (
	countrydto "dewetour/dto/country"
	dto "dewetour/dto/result"
	"dewetour/models"
	"dewetour/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Countryhand struct {
	CountryRepository repositories.CountryRepository
}

func HandlerCountry(CountryRepository repositories.CountryRepository) *Countryhand {
	return &Countryhand{CountryRepository}
}

func (h *Countryhand) FindCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	country, err := h.CountryRepository.FindCountry()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: country}
	json.NewEncoder(w).Encode(response)
}
func (h *Countryhand) GetCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: country}
	json.NewEncoder(w).Encode(response)
}
func convertCountryResponse(u models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		// ID:       u.ID,
		Name: u.Name,
	}
}
func convertCountryDeleteResponse(u models.Country) countrydto.CountryDeleteResponse {
	return countrydto.CountryDeleteResponse{
		ID: u.ID,
		// Name: u.Name,
	}
}

func (h *Countryhand) CreateCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(UserInfo["id"].(float64))

	request := countrydto.CreateCountryRequest{
		Name: r.FormValue("name"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	country := models.Country{
		Name:   request.Name,
		UserID: userID,
	}

	data, err := h.CountryRepository.CreateCountry(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *Countryhand) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userID := int(UserInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	country, err := h.CountryRepository.GetCountry(int(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.FormValue("name") != "" {
		country.Name = r.FormValue("name")
	}

	data, err := h.CountryRepository.UpdateCountry(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertCountryResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *Countryhand) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CountryRepository.DeleteCountry(country, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertCountryDeleteResponse(data)}
	json.NewEncoder(w).Encode(response)
}
