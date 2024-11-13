package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/faustcelaj/social_project/internal/store"
	"github.com/go-chi/chi/v5"
)

type userkey string

const userCtx userkey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUsersFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// middlewear to fetch userID
func (app *application) usersContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "userID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.InternalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetById(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.NotFoundResponse(w, r, err)
			default:
				app.InternalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUsersFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
