package handlers

import (
	dto "dewetour/dto/result"
	transactiondto "dewetour/dto/transaction"
	"dewetour/models"
	"dewetour/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type Transactionhand struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *Transactionhand {
	return &Transactionhand{TransactionRepository}
}

func (h *Transactionhand) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *Transactionhand) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	// trip.Image = path_file + trip.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

// var file = "http://localhost:9000/uploads/"
func (h *Transactionhand) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	UserID := int(UserInfo["id"].(float64))

	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

	request := new(transactiondto.CreateTransaction)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Println(request)

	var TransIdIsMatch = false
	var TransactionId int
	for !TransIdIsMatch {
		TransactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(TransactionId)
		if transactionData.ID == 0 {
			TransIdIsMatch = true
		}
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TripID:     request.TripID,
		UserID:     UserID,
		// Attachment: filename,
	}
	fmt.Println(transaction)

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	fmt.Println(data)

	get, err := h.TransactionRepository.GetTransaction(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	fmt.Println(get)
	var s = snap.Client{}
	s.New("SB-Mid-server-6usfxF-FXVlMxthSYolOdn4H", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(get.ID),
			GrossAmt: int64(get.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: get.User.FullName,
			Email: get.User.Email,
		},
	}
	fmt.Println(req)

	snapResp, _ := s.CreateTransaction(req)

	fmt.Println(snapResp)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)

}

// 	transaction.Attachment = path_file + transaction.Attachment
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: get}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *Transactionhand) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	userID := int(UserInfo["id"].(float64))

// 	dataContex := r.Context().Value("dataFile")
// 	filename := dataContex.(string)

// 	req := transactiondto.UpdateTransaction{
// 		Attachment: filename,
// 		UserID:     userID,
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	getData, err := h.TransactionRepository.GetTransaction(int(id))

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	qty, _ := strconv.Atoi(r.FormValue("qty"))
// 	if qty != 0 {
// 		getData.CounterQty = qty
// 	}
// 	total, _ := strconv.Atoi(r.FormValue("total"))
// 	if total != 0 {
// 		getData.Total = total
// 	}
// 	if r.FormValue("status") != "" {
// 		getData.Status = r.FormValue("status")
// 	}
// 	tripID, _ := strconv.Atoi(r.FormValue("trip_id"))
// 	if tripID != 0 {
// 		getData.TripID = tripID
// 	}
// 	if req.Attachment != "" {
// 		getData.Attachment = req.Attachment
// 	}

// 	// update data
// 	data, err := h.TransactionRepository.UpdateTransaction(getData)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// get data
// 	GetID, err := h.TransactionRepository.GetTransaction(data.ID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// 	// success
// 	getData.Attachment = path_file + getData.Attachment
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: GetID}
// 	json.NewEncoder(w).Encode(response)
// }

// Notifikasi
func (h *Transactionhand) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	// convOrderId, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetOneTransaction(orderId) // fmt.Println(transactionStatus, fraudStatus, orderId, transaction)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", transaction.ID)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", transaction.ID)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", transaction.ID)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateTransaction("failed", transaction.ID)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateTransaction("failed", transaction.ID)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", transaction.ID)
	}

	w.WriteHeader(http.StatusOK)
}

func SendMail(status string, transaction models.Transaction) {

	// var s = snap.Client{}
	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = "afriandifrans@gmail.com"
		var CONFIG_AUTH_PASSWORD = "xhsjiljtwydfmjeh"

		var tripName = transaction.Trip.Title
		var price = strconv.Itoa(transaction.Trip.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : % s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, tripName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}
