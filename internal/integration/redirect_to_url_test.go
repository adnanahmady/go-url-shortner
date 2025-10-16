package integration

import (
	"net/http"
	"testing"

	"github.com/adnanahmady/go-url-shortner/pkg/test"
	"github.com/adnanahmady/go-url-shortner/pkg/test/assert"
)

func TestRedirectToUrl(t *testing.T) {
	app, err := test.Setup()
	assert.NoError(t, err)

	app.StoreManager.Lock()
	app.StoreManager.Set("123456", "https://www.google.com")
	app.StoreManager.Unlock()

	t.Run("given url when exists then should redirect to original url", func(t *testing.T) {
		path := "/123456"
		handler := app.V1Handlers.RedirectToOriginalUrl

		rec, err := test.Get(app, path, handler)
		assert.NoError(t, err)

		assert.Equal(t, rec.Code, http.StatusFound)
		assert.Equal(t, rec.Header().Get("Location"), "https://www.google.com")
	})

	t.Run("given url when not exists then should return not found", func(t *testing.T) {
		path := "/not-exists"
		handler := app.V1Handlers.RedirectToOriginalUrl

		rec, err := test.Get(app, path, handler)
		assert.NoError(t, err)

		assert.Equal(t, rec.Code, http.StatusNotFound)
	})
}
