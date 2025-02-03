package server

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	_ "time"

	"github.com/arabenjamin/fetch/app"
)

/*
func logger(thisLogger *log.Logger) Middleware {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				thisLogger.Println(r.URL.Path, time.Now().Unix())
			}()
			next(w, r)
		}
	}
}
*/

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
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

func respond(res http.ResponseWriter, payload map[string]interface{}) {

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
	}

	/*Validate Payload */
	if err := json.NewDecoder(req.Body).Decode(&reciept); err != nil {
		http.Error(resp, "Invalid request body", http.StatusBadRequest)
		return
	}

	/*SaveReciept Generates an Id */
	r, err := app.SaveReciept(reciept)
	if err != nil {
		http.Error(resp, "Error saving reciept", http.StatusInternalServerError)
	}

	/* Response payload */
	payload := map[string]interface{}{
		"id": r.Id,
	}

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
	// TODO: Where do we look up this id ?
	id := req.PathValue("id")
	fmt.Printf("ID: %v\n", id)
	points, _ := app.GetRecieptByID(id)
	/*if err != nil {
		HandleError(resp, err, http.StatusNotFound)
		//fmt.Println(err)

		//http.Error(resp, "Reciept not found", http.StatusNotFound)
	}*/
	/* Return the points awarded to given*/
	payload := map[string]interface{}{
		"points": points,
	}

	respond(resp, payload)
}

/* For Testing and Troubleshooting purposes */
func ping(resp http.ResponseWriter, req *http.Request) {

	/* Response payload */
	payload := map[string]interface{}{
		"status":  "ok",
		"message": "pong!",
	}

	/* Ping Pong */
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
