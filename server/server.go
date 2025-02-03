package server

import (
  
  _"fmt"
  _"io"
  _"log"
  _"crypto/md5"
	"encoding/json"
	"net/http"
  "github.com/google/uuid"
	_"time"
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

func respond(res http.ResponseWriter, payload map[string]interface{}) {

	resp_json, _ := json.Marshal(payload)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(http.StatusOK)
	res.Write(resp_json)

}

/*  Submits a receipt for processing */
func processReciept(resp http.ResponseWriter, req *http.Request){


    /* Validate request method */
    if req.Method != http.MethodPost {
      http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
    }

    /*Validate Payload */
    var requestData struct {
			Reciept bool `json:"reciept"`
		}
  
    if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			http.Error(resp, "Invalid request body", http.StatusBadRequest)
			return
		}

    /*Generate and Store new ID */
    id := uuid.New()
    payload := map[string]interface{}{
      "id" : id.String()
    }


    /* Return the new id */
    respond(resp, payload) 
}

func getRecieptById(resp http.ResponseWriter, req *http.Request){


    /* Validate request method */
    if req.Method != http.MethodGet{
      http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
    }

    /*Make sure we have that id*/
    // TODO: Where do we look up this id ?    

    /* Return the points awarded to given*/
    paylaod := map[string]interface{

      "points":42
    }

    respond(resp, payload)
}

/* For Testing and Troubleshooting purposes */
func ping(resp http.ResponseWriter, req *http.Request){

    /* Response payload */
  	resp_payload := map[string]interface{}{
		  "status":       "ok",
		  "message":      "pong!",
	  }

    /* Ping Pong */
    respod(resp, payload)
}

func StartServer() error {
  
  /*Run our server or return an error */
  mux := http.NewServeMux()
  
  mux.HandleFunc("/ping", ping)   	
  mux.HandleFunc("/receipts/process", processReciept)
  mux.HandleFunc("/receipts/{id}/points", getRecieptById)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}
	return nil

}



