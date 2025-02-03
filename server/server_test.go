package server

import (

	//"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/arabenjamin/fetch/server"
)

func TestServerPing(t *testing.T) {

	t.Run("Returns Pong", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/ping", nil)
		res := httptest.NewRecorder()
		ping(res, req)

		//server := server.StartServer()
		resp := res.Result()
		expected := 200

		if resp.StatusCode != expected {
			t.Errorf("Recieved %q, instead got %q", resp.StatusCode, expected)
		}
	})
}

func TestServerPostReciept(t *testing.T) {

	t.Run("Tests Process Reciepts Endpoint", func(t *testing.T) {

		req := httptest.NewRequest("POST", "/reciepts/process", nil)
		res := httptest.NewRecorder()

		SaveAndProcessReciept(res, req)

		resp := res.Result()
		expected := 201

		if resp.StatusCode != expected {
			t.Errorf("Expecting %q, Instead got %q", expected, resp.StatusCode)
		}

	})
}

func TestServerGetPointsById(t *testing.T) {

	t.Run("Tests GetPointsByID", func(t *testing.T) {

		req := httptest.NewRequest("GET", "receipts/{id}/points", nil)
		res := httptest.NewRecorder()

		GetRecieptById(res, req)
		resp := res.Result()
		expected := 200

		if resp.StatusCode != expected {
			t.Errorf("Expecting %q, Instead recieved %q", expected, resp.StatusCode)
		}

		/* Test to be sure that we returned the right pointas*/

	})
}
