package app

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

var Reciepts []Reciept

type Reciept struct {
	Id           string
	Points       string `json:"points"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDesription string
	Price           string
}

func GetRecieptByID(id string) (Reciept, error) {

	fmt.Printf("Reciept ID: %v\n", id)
	log.Printf("Reciept ID: %v\n", id)
	/*I need to find the reciept in the Reciepts slice*/
	var r Reciept

	if len(Reciepts) == 0 {
		return r, errors.New("No Reciepts found, add reciept first")
	}

	for _, reciept := range Reciepts {

		if reciept.Id == id {
			r = reciept
		}

	}

	/*
		if r.Id == "" {

			return r, errors.New("Reciept not found by id")
		}*/

	return r, nil

}

func SaveReciept(r Reciept) (Reciept, error) {

	/*Create a uuid for the id of the reciept*/
	id := uuid.New().String()
	r.Id = id

	/*Process the reciept*/
	r, err := ProcessReciept(r)
	if err != nil {
		return r, err
	}

	/*
	  I need some non persistant way to save the recipet,
	  for now I will just add it to the Reciepts slice
	*/
	Reciepts = append(Reciepts, r)

	return r, nil
}

func ProcessReciept(r Reciept) (Reciept, error) {

	/*I need to calculate the points for the reciept*/
	r.Points = "42"

	return r, nil
}
