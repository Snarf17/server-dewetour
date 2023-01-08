package handlers

import (
	dto "dewetour/dto/result"
	tripdto "dewetour/dto/trip"
	"dewetour/models"
	"dewetour/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Tripshand struct {
	TripRepository repositories.TripRepository
}

func HandlerTrips(TripRepository repositories.TripRepository) *Tripshand {
	return &Tripshand{TripRepository}
}

func (h *Tripshand) FindTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trips, err := h.TripRepository.FindTrip()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	for i, p := range trips {
		trips[i].Image = path_file + p.Image
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}
func (h *Tripshand) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	trip.Image = path_file + trip.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

func convertTripResponse(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		// ID:       u.ID,
		Title:          u.Title,
		CountryID:      u.CountryID,
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
		Image:          u.Image,
		UserID:         u.UserID,
	}
}

// func convertIdResponse(u models.Trip) tripdto.CountryIdResponse {
// 	return tripdto.CountryIdResponse{
// 		Title:          u.Title,
// 		Accomodation:   u.Accomodation,
// 		Country:        u.Country,
// 		Transportation: u.Title,
// 		Eat:            u.Eat,
// 		Day:            u.Day,
// 		Night:          u.Night,
// 		DateTrip:       u.DateTrip,
// 		Price:          u.Price,
// 		Description:    u.Description,
// 		Image:          u.Image,
// 	}
// }

var path_file = "http://localhost:9000/uploads/"

func (h *Tripshand) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(UserInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// request := new(tripdto.CreateTripRequest)

	price, _ := strconv.Atoi(r.FormValue("price"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))
	// DateTrip, _ := strconv.Atoi(r.FormValue("date_trip"))
	CountryID, _ := strconv.Atoi(r.FormValue("countryID"))
	request := tripdto.CreateTripRequest{
		Title:          r.FormValue("title"),
		CountryID:      CountryID,
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transport"),
		Eat:            r.FormValue("eat"),
		Day:            day,
		Night:          night,
		Price:          price,
		Quota:          quota,
		Description:    r.FormValue("desc"),
		Image:          filename,
		// DateTrip:       r.FormValue("edate"),
	}
	fmt.Println(night)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trip := models.Trip{
		Title:          request.Title,
		CountryID:      request.CountryID,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          request.Image,
		UserID:         userID,
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	get, err := h.TripRepository.GetTrip(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	trip.Image = path_file + trip.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: get}
	json.NewEncoder(w).Encode(response)
}
func (h *Tripshand) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(UserInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	req := tripdto.UpdateTripRequest{
		Image:  filename,
		UserID: userID,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.GetTrip(int(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.FormValue("title") != "" {
		trip.Title = r.FormValue("title")
	}
	country_idInput, _ := strconv.Atoi(r.FormValue("country_id"))
	if country_idInput != 0 {
		trip.CountryID = country_idInput
	}
	if r.FormValue("accomodation") != "" {
		trip.Accomodation = r.FormValue("accomodation")
	}
	if r.FormValue("transportation") != "" {
		trip.Transportation = r.FormValue("transport")
	}
	if r.FormValue("eat") != "" {
		trip.Eat = r.FormValue("eat")
	}
	day, _ := strconv.Atoi(r.FormValue("day"))
	if day != 0 {
		trip.Day = day
	}
	night, _ := strconv.Atoi(r.FormValue("night"))
	if night != 0 {
		trip.Night = night
	}
	date_trip, _ := time.Parse("2006-01-02", r.FormValue("date_trip"))
	if date_trip.IsZero() {
		date_trip := trip.DateTrip
		trip.DateTrip = date_trip
	}
	priceInput, _ := strconv.Atoi(r.FormValue("price"))
	if priceInput != 0 {
		trip.Price = priceInput
	}
	quotaInput, _ := strconv.Atoi(r.FormValue("quota"))
	if quotaInput != 0 {
		trip.Quota = quotaInput
	}
	if r.FormValue("description") != "" {
		trip.Description = r.FormValue("desc")
	}
	if req.Image != "" {
		trip.Image = req.Image
	}

	// result = make([]models.Trip, len(filename))

	// fmt.Println(dataContex)

	// fmt.Println(trip.ID)
	// fmt.Println(trip.CountryID)
	// update data
	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// get data
	test, err := h.TripRepository.GetTrip(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: test}
	json.NewEncoder(w).Encode(response)
}
func convertDeleteTripResponse(u models.Trip) tripdto.TripDeleteResponse {
	return tripdto.TripDeleteResponse{
		ID: u.ID,
	}
}

func (h *Tripshand) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.TripRepository.GetTrip(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(user, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Data: convertDeleteTripResponse(data)}
	json.NewEncoder(w).Encode(response)
}
