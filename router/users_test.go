package router

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kk-no/testable-api/store"
)

func NewMockUser() *User { return &User{Store: store.NewMockUser()} }

func TestUser_Create(t *testing.T) {
	const method = http.MethodPost

	u := NewMockUser()

	type want struct {
		code int
		body string
	}

	tests := []struct {
		name   string
		target string
		want   want
	}{
		{
			name:   "success",
			target: "/users",
			want: want{
				code: http.StatusCreated,
				body: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(method, tt.target, nil)
			u.Create(w, r)

			if w.Code != tt.want.code {
				t.Errorf("unexpected status code: want %v got %v", tt.want.code, w.Code)
			}
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	const method = http.MethodGet

	u := NewMockUser()

	tests := []struct {
		name   string
		target string
	}{
		{
			name:   "success",
			target: "/users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(method, tt.target, nil)
			u.GetByID(w, r)

			log.Println(w.Code)
		})
	}
}

func TestUser_Update(t *testing.T) {
	const method = http.MethodPatch

	u := NewMockUser()

	tests := []struct {
		name   string
		target string
	}{
		{
			name:   "success",
			target: "/users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(method, tt.target, nil)
			u.GetByID(w, r)

			log.Println(w.Code)
		})
	}
}
