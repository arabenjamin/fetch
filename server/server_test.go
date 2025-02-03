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

	t.Run("Returns_200", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/ping", nil)
		res := httptest.NewRecorder()

		ping(res, req)

		resp := res.Result()
		expected := 200

		if resp.StatusCode != expected {
			t.Errorf("Recieved %d, instead got %d", resp.StatusCode, expected)
		}

	})
	t.Run("Returns_Pong", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/ping", nil)
		res := httptest.NewRecorder()

		ping(res, req)

		resp := res.Result()
		defer resp.Body.Close()

		var data RecieptResponse

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			t.Error(err)
		}

		if data.Message != "pong!" {
			t.Errorf("Expected pong, got %v", data.Message)
		}
	})

}

func TestServerPostReciept(t *testing.T) {

	t.Run("Tests_returns_201", func(t *testing.T) {

		requestBody := `{ "retailer": "Walmart", "purchaseDate": "2020-01-01", "purchaseTime": "12:00", "items": [ { "shortDescription": "item1", "price": "1.00" }, { "shortDescription": "item2", "price": "2.00" } ], "total": "3.00" }`
		req, err := http.NewRequest("POST", "/reciepts/process", strings.NewReader(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		SaveAndProcessReciept(res, req)
		// Check the response status code
		if status := res.Code; status != http.StatusCreated {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

	})

	// Test for Bad Reciepts
	t.Run("Tests_for_empty_ticket", func(t *testing.T) {

		requestBody := `{}`
		req, err := http.NewRequest("POST", "/reciepts/process", strings.NewReader(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		SaveAndProcessReciept(res, req)
		if status := res.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

	})
	// Test for	400
	t.Run("Tests_for_bad_ticket", func(t *testing.T) {
		requestBody := `{ "retailer": "" }`
		req, err := http.NewRequest("POST", "/reciepts/process", strings.NewReader(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		SaveAndProcessReciept(res, req)
		// Check the response status code
		if status := res.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	// Test for	404
	/*
		t.Run("Tests_for_bad_endpoint_url_returns_404", func(t *testing.T) {
			//requestBody := `{}`
			requestBody := `{ "retailer": "Walmart", "purchaseDate": "2020-01-01", "purchaseTime": "12:00", "items": [ { "shortDescription": "item1", "price": "1.00" }, { "shortDescription": "item2", "price": "2.00" } ], "total": "3.00" }`
			req, err := http.NewRequest("POST", "/BadEndpointUrl/", strings.NewReader(requestBody))
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("Requested URL: %v", req.URL)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()

			SaveAndProcessReciept(res, req)

			// Check the response status code
			if status := res.Code; status != http.StatusNotFound {
				//url_used := res.Result().Request.URL.String()
				t.Logf("Requested URL: %v ", req.URL)
				t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
			}
		})
	*/
	t.Run("Tests_returns_ID", func(t *testing.T) {

		requestBody := `{ "retailer": "Walmart", "purchaseDate": "2020-01-01", "purchaseTime": "12:00", "items": [ { "shortDescription": "item1", "price": "1.00" }, { "shortDescription": "item2", "price": "2.00" } ], "total": "3.00" }`
		req, err := http.NewRequest("POST", "/reciepts/process", strings.NewReader(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		SaveAndProcessReciept(res, req)
		var data RecieptResponse
		// Check that we got an ID back in the json response
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			t.Error(err)
		}
		if data.Id == "" {
			t.Errorf("Expected ID, got %v", data.Id)
		}

	})

}

func TestServerGetPointsById(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/receipts/{id}/points", GetRecieptById)

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	/*Clear the global list of Reciepts */
	app.Reciepts = nil
	test_reciept := app.Reciept{Id: "foo", Points: "100", Retailer: "Walmart", PurchaseDate: "2020-01-01", PurchaseTime: "12:00", Items: []app.Item{{ShortDesription: "item1", Price: "1.00"}, {ShortDesription: "item2", Price: "2.00"}}, Total: "3.00"}
	test_recipet_points, err := app.ProcessReciept(test_reciept)
	if err != nil {
		t.Fatal(err)
	}
	app.Reciepts = append(app.Reciepts, test_recipet_points)

	/* Test returns 200*/
	t.Run("Tests_returns_200", func(t *testing.T) {

		curennt_url := testServer.URL + "/receipts/foo/points"
		req, err := http.NewRequest("GET", curennt_url, nil)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

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

	})

	/* Test returns 404*/
	t.Run("Test_returns_404_missing_id", func(t *testing.T) {

		mux := http.NewServeMux()
		mux.HandleFunc("/receipts/{id}/points", GetRecieptById)

		testServer := httptest.NewServer(mux)
		defer testServer.Close()

		curennt_url := testServer.URL + "/receipts/Batman/points"
		req, err := http.NewRequest("GET", curennt_url, nil)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		expected := 404

		if resp.StatusCode != expected {
			t.Errorf("Expecting %d, Instead recieved %d", expected, resp.StatusCode)
		}

	})

	/* Test to be sure that we returned the right points*/
	t.Run("Test_return_points", func(t *testing.T) {

		curennt_url := testServer.URL + "/receipts/foo/points"
		req, err := http.NewRequest("GET", curennt_url, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		var data app.Reciept
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			t.Errorf("Expected points, got error: %v", err)
		}

		if data.Points == "" {
			t.Errorf("Expected %v points, got nothing", test_recipet_points)
		}

		if data.Points != test_recipet_points.Points {
			t.Errorf("ProcessingReciept is borked, Points returned %v", data.Points)
		}

	})

}
