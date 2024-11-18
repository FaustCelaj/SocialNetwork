package main

import (
	"net/http"
)

// pagination, search, filters

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(22))
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
