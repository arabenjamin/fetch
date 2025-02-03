package app

import (

	//"fmt"
	"errors"

	"github.com/google/uuid"
)

var reciepts []Reciept

type Reciept struct {
	Id           string
	Points       int
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	shortDesription string
	price           string
}

func GetRecieptByID(id string) (Reciept, error) {

	/*I need to find the reciept in the reciepts slice*/
	var r Reciept

	if len(reciepts) == 0 {
		return r, errors.New("No reciepts found, add recipet first")
	}

	for _, reciept := range reciepts {

		if reciept.Id == id {
			r = reciept
		}

	}

	if r.Id == "" {

		return r, errors.New("Reciept not found by id")
	}

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
	  for now I will just add it to the reciepts slice
	*/
	reciepts = append(reciepts, r)

	return r, nil
}

func ProcessReciept(r Reciept) (Reciept, error) {

	/*I need to calculate the points for the reciept*/
	r.Points = 42

	return r, nil
}
