package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGetBlocks(t *testing.T) {
	req, err := http.NewRequest("GET", "/blocks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGetBlocks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}

	expected := `[{"index":0,"timestamp":`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Expected body to contain %v, got %v", expected, rr.Body.String())
	}
}

func TestHandleMineBlock(t *testing.T) {
	payload := strings.NewReader(`{"data":"Test Block"}`)
	req, err := http.NewRequest("POST", "/mine", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleMineBlock)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status %v, got %v", http.StatusCreated, status)
	}

	expected := `"data":"Test Block"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Expected body to contain %v, got %v", expected, rr.Body.String())
	}
}