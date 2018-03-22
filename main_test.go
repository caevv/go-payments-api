package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestEmptyTable(t *testing.T) {
	req, _ := http.NewRequest("POST", "/payment", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != `{"id":1}` {
		t.Errorf("Expected id. Got %s", body)
	}
}

func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	CreatePayment(w, r)

	return w
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
