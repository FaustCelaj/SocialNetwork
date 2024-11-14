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

// get a single user
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUsersFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

// follow a user
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	// The user we want to follow
	followUser := getUsersFromCtx(r)

	// TODO: revert back to auth
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, followUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}

	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// unfollow a user
func (app *application) unFollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUsersFromCtx(r)

	// TODO: revert back to auth
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
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
				app.notFoundResponse(w, r, err)
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
