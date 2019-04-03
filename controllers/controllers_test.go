package controllers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//testcase for the success scenario
func TestCrawlURLS(t *testing.T) {
	req, err := http.NewRequest("GET", "/crawl?url=https://access.redhat.com/support/contact", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CrawlURLS)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := string(`{"https://access.redhat.com/support/contact":true,
	"https://access.redhat.com/support/contact/":true,"https://access.redhat.com/support/contact/Sales/":true,
	"https://access.redhat.com/support/contact/customerService/":true,
	"https://access.redhat.com/support/contact/technicalSupport/":true}`)
	fmt.Println("rr.Body", rr.Body.String())
	assert.Equal(t, rr.Code, http.StatusOK)

	assert.JSONEq(t, expectedOutput, rr.Body.String())

}

//testcase for the bad request status code if you forget to send url in query param
func TestCrawlURLSBadRequest(t *testing.T) {

	req, err := http.NewRequest("GET", "/crawl", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CrawlURLS)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
	assert.Equal(t, rr.Code, http.StatusBadRequest)

}
