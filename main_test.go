package main

import (
	"testing"

	"github.com/adnanahmady/go-url-shortner/pkg/test"
	"github.com/adnanahmady/go-url-shortner/pkg/test/assert"
)

func TestLoadAndSaveUrls(t *testing.T) {
	// Arrange
	app, err := test.Setup()
	assert.NoError(t, err)

	app.StoreManager.Lock()
	app.StoreManager.Set("123456", "https://www.google.com")
	app.StoreManager.Set("123457", "https://www.example.com")
	app.StoreManager.Unlock()

	saveUrls(app)

	app.StoreManager.Lock()
	app.StoreManager.Clear()
	app.StoreManager.Unlock()

	// Act
	loadUrls(app)

	// Assert
	app.StoreManager.Lock()
	defer app.StoreManager.Unlock()

	assert.Truef(t, app.StoreManager.Has("123456"), "failed to assert 123456 is in store")
	assert.Truef(t, app.StoreManager.Has("123457"), "failed to assert 123457 is in store")
	assert.Equal(t, app.StoreManager.Count(), 2)
}
