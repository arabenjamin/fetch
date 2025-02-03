package server

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/arabenjamin/fetch/app"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type RecieptResponse struct {
	Id           string `json:"id,omitempty"`
	Points       int    `json:"points,omitempty"`
	Status       string `json:"status,omitempty"`
	Message      string `json:"messgage,omitempty"`
	Errormessage string `json:"errormessage,omitempty"`
}

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Message: err.Error(),
		Code:    statusCode, // Optional
	}

	//json.NewEncoder(w).Encode(errorResponse)
	resp_json, _ := json.Marshal(errorResponse)
	w.Write(resp_json)
}

func respond(res http.ResponseWriter, payload RecieptResponse) {

	resp_json, _ := json.Marshal(payload)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(resp_json)

}

/*  Submits a receipt for processing */
func SaveAndProcessReciept(resp http.ResponseWriter, req *http.Request) {

	reciept := app.Reciept{}
	/* Validate request method */
	if req.Method != http.MethodPost {
		http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
		//HandleError(resp, err, http.StatusNotFound)
		return
	}

	content_type := req.Header.Get("Content-Type")
	if !strings.Contains(content_type, "application/json") {
		http.Error(resp, "Incorrect content-type", http.StatusBadRequest)
		return
	}

	/*Dump it into json */
	if err := json.NewDecoder(req.Body).Decode(&reciept); err != nil {
		http.Error(resp, "Invalid request body", http.StatusBadRequest)
		return
	}

	/*Check to make sure all fields are present*/
	validate := validator.New()
	err := validate.Struct(reciept)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(resp, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	/* Check if at least one Item is present*/
	if len(reciept.Items) < 1 {
		http.Error(resp, "Invalid request body", http.StatusBadRequest)
	}

	/*SaveReciept Generates an Id */
	r, err := app.SaveReciept(reciept)
	if err != nil {

		http.Error(resp, "Error saving reciept", http.StatusInternalServerError)
	}

	/* Response payload */
	var payload RecieptResponse
	payload.Id = r.Id

	/* Return the new id */
	resp.WriteHeader(http.StatusCreated)
	respond(resp, payload)
}

func GetRecieptById(resp http.ResponseWriter, req *http.Request) {

	/* Validate request method */
	if req.Method != http.MethodGet {
		http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
	}

	/*Make sure we have that id*/
	id := req.PathValue("id")
	reciept, err := app.GetRecieptByID(id)
	if err != nil {
		HandleError(resp, err, http.StatusNotFound)
		//http.Error(resp, "Reciept not found", http.StatusNotFound)
	}

	/* Return the points awarded to given*/
	var payload RecieptResponse
	payload.Points, _ = strconv.Atoi(reciept.Points)

	respond(resp, payload)
}

/* For Testing and Troubleshooting purposes */
func ping(resp http.ResponseWriter, req *http.Request) {

	/* Response payload */

	var payload RecieptResponse
	/* Ping Pong */
	payload.Message = "pong!"
	payload.Status = "ok"
	resp.WriteHeader(http.StatusOK)
	respond(resp, payload)
}

func StartServer() error {

	/*Run our server or return an error */
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/receipts/process", SaveAndProcessReciept)
	mux.HandleFunc("/receipts/{id}/points", GetRecieptById)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}
	return nil

}
