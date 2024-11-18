package main

import (
	"net/http"

	"github.com/faustcelaj/social_project/internal/store"
)

// pagination, search, filters

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(22), fq)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}