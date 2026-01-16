package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllPages(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		handler http.HandlerFunc
	}{
		{"Home", "/", homePage},
		{"Courses", "/courses", coursePage},
		{"About", "/about", aboutPage},
		{"Contact", "/contact", contactPage},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			tt.handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("expected 200, got %d", rr.Code)
			}
		})
	}
}
