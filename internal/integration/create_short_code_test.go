package integration

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/adnanahmady/go-url-shortner/pkg/test"
	"github.com/adnanahmady/go-url-shortner/pkg/test/assert"
)

func TestCreateShortCode(t *testing.T) {
	app, err := test.Setup()
	assert.NoError(t, err)

	t.Run("given url when shorten api called then should store shorten url in memory", func(t *testing.T) {
		testUrl := "https://www.google.com"
		data := url.Values{}
		data.Set("url", testUrl)

		rec, err := test.Post(app, "/shorten", data, app.V1Handlers.CreateShortUrl)
		assert.NoError(t, err)

		assert.Equal(t, rec.Code, http.StatusOK)
		app.StoreManager.Lock()
		defer app.StoreManager.Unlock()
		_, has := app.StoreManager.HasUrl(testUrl)
		assert.Truef(t, has, "short url should be created")
	})

	t.Run("given url when shorten then short url should return", func(t *testing.T) {
		testUrl := "https://www.google2.com"
		data := url.Values{}
		data.Set("url", testUrl)

		rec, err := test.Post(app, "/shorten", data, app.V1Handlers.CreateShortUrl)
		assert.NoError(t, err)

		app.StoreManager.Lock()
		defer app.StoreManager.Unlock()
		shortCode, _ := app.StoreManager.HasUrl(testUrl)
		assert.Equal(t, rec.Body.String(), fmt.Sprintf("http://localhost:5000/%v\n", shortCode))
	})

	t.Run("given url when its missing then should return bad request", func(t *testing.T) {
		initialCount := app.StoreManager.Count()
		data := url.Values{}

		rec, err := test.Post(app, "/shorten", data, app.V1Handlers.CreateShortUrl)
		assert.NoError(t, err)

		app.StoreManager.Lock()
		defer app.StoreManager.Unlock()
		assert.Equal(t, rec.Code, http.StatusBadRequest)
		assert.Equal(t, rec.Body.String(), "URL is required")
		assert.Equal(t, app.StoreManager.Count()-initialCount, 0)
	})
}
