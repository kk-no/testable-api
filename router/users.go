package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kk-no/testable-api/models"
	"github.com/kk-no/testable-api/store"
)

type User struct{ Store *store.User }

func NewUser() *User {
	return &User{store.NewUser()}
}

func Users(r chi.Router) {
	u := NewUser()
	r.Get("/{userID}", u.GetByID)
	r.Post("/", u.Create)
	r.Patch("/{userID}", u.Update)
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := u.Store.FindByID(ctx, chi.URLParam(r, "userID"))
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := u.Store.Create(ctx, &models.User{ID: "1", Name: "tester"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := u.Store.Update(ctx, &models.User{ID: "1", Name: "updater"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
