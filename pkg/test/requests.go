package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/adnanahmady/go-url-shortner/internal"
	request "github.com/adnanahmady/go-url-shortner/pkg/reqeust"
)

func Post(app *internal.App, path string, body url.Values, handler http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	ctx := request.SetLogger(req.Context(), app.Logger)
	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	return rec, nil
}

func Get(app *internal.App, path string, handler http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	ctx := request.SetLogger(req.Context(), app.Logger)
	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	return rec, nil
}
