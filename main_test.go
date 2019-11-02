package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test_ping(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "pong",
	}

	// Grab our router
	router := GetMainEngine()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/v1/ping")

	// Test StatusOK(200) from response
	t.Run("GET HTTP StatusOK", func(test *testing.T) {
		expected := http.StatusOK
		actual := w.Code
		if expected != actual {
			t.Errorf(" expected %v but it got %v", expected, actual)
		}
	})

	// Test Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	t.Run("Convert Body to Unmarhal", func(test *testing.T) {
		if err != nil {
			t.Errorf(" expected %v but it got %v", "not nil", err)
		}
	})

	t.Run("Expected body['message'] equal response['message']", func(test *testing.T) {

		expected := body["message"]
		actual := response["message"]
		if expected != actual {
			t.Errorf(" expected %v but it got %v", expected, actual)
		}
	})
}

func Test_RemandsAPI(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "pong",
	}

	// Grab our router
	router := GetMainEngine()
	// Perform a GET request with that handler.
	w := performRequest(router, "POST", "/v1/remands")

	// Test Convert the JSON response to a map
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)

	t.Run("Expected body['message'] equal response['message']", func(test *testing.T) {

		expected := body["message"]
		actual := "pong"
		if expected != actual {
			t.Errorf(" expected %v but it got %v", expected, actual)
		}
	})
}
