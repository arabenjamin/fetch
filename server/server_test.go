package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arabenjamin/fetch/app"
)

func TestServerPing(t *testing.T) {

	t.Run("Returns 200", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/ping", nil)
		res := httptest.NewRecorder()

		ping(res, req)

		resp := res.Result()
		expected := 200

		if resp.StatusCode != expected {
			t.Errorf("Recieved %d, instead got %d", resp.StatusCode, expected)
		}

	})
	t.Run("Returns Pong", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/ping", nil)
		res := httptest.NewRecorder()

		ping(res, req)

		resp := res.Result()
		defer resp.Body.Close()

		var data map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			t.Error(err)
		}

		if data["message"] != "pong!" {
			t.Errorf("Expected pong, got %v", data["message"])
		}
	})

}

func TestServerPostReciept(t *testing.T) {

	t.Run("Tests ProcessReciepts 200", func(t *testing.T) {

		requestBody := `{ "retailer": "Walmart", "purchaseDate": "2020-01-01", "purchaseTime": "12:00", "items": [ { "shortDescription": "item1", "price": "1.00" }, { "shortDescription": "item2", "price": "2.00" } ], "total": "3.00" }`
		req, err := http.NewRequest("POST", "/reciepts/process", strings.NewReader(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()

		SaveAndProcessReciept(res, req)

		// Check the response status code
		if status := res.Code; status != http.StatusCreated {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

	})

	// TODO: Add a test to check the response body and ensure that the id is returned

}

func TestServerGetPointsById(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/receipts/{id}/points", GetRecieptById)

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	/*Clear the global list of Reciepts */
	app.Reciepts = nil
	test_reciept := app.Reciept{Id: "foo", Points: "100", Retailer: "Walmart", PurchaseDate: "2020-01-01", PurchaseTime: "12:00", Items: []app.Item{{ShortDesription: "item1", Price: "1.00"}, {ShortDesription: "item2", Price: "2.00"}}, Total: "3.00"}
	app.Reciepts = append(app.Reciepts, test_reciept)

	/*
		  for _, reciept := range app.Reciepts {
				t.Logf("Reciept ID: %v\n", reciept.Id)
				if reciept.Id == "foo" {
					t.Logf("Reciept Found: %v\n", reciept)
				}
			}*/

	t.Run("Tests_GetPointsByID_returns_200", func(t *testing.T) {

		curennt_url := testServer.URL + "/receipts/foo/points"
		req, err := http.NewRequest("GET", curennt_url, nil)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		t.Logf("GET Url used: %v", req.URL)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		/*data := ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			t.Error(err)
		}*/

		expected := 200

		if resp.StatusCode != expected {
			t.Errorf("Expecting %d, Instead recieved %d", expected, resp.StatusCode)
		}

		/* Test to be sure that we returned the right pointas*/

	})
}
