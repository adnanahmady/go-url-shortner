package presentation

import (
	"net/http"

	"github.com/adnanahmady/go-url-shortner/internal/application"
	request "github.com/adnanahmady/go-url-shortner/pkg/reqeust"
)

type V1Handlers struct {
	create application.CreateShortUrlUseCase
	get application.GetShortUrlUseCase
}

func NewV1Handlers(
	create application.CreateShortUrlUseCase,
	get application.GetShortUrlUseCase,
) *V1Handlers {
	return &V1Handlers{
		create: create,
		get: get,
	}
}

func (h *V1Handlers) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := request.GetLogger(ctx).Section("V1Handlers", "CreateShortUrl")

	if r.Method != http.MethodPost {
		logger.Errorf("Method (%v) is not allowed", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	url := r.FormValue("url")
	if url == "" {
		logger.Errorf("URL is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("URL is required"))
		return
	}
	logger.Infof("Received url to shorten (%v)", url)

	shortUrl, err := h.create.Create(ctx, url)
	if err != nil {
		logger.Errorf("Error creating short url (%v)", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	logger.Infof("Short url created (%v)", shortUrl)
	w.Write([]byte(shortUrl))
}

func (h *V1Handlers) RedirectToOriginalUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := request.GetLogger(ctx).Section("V1Handlers", "RedirectToOriginalUrl")

	if r.Method != http.MethodGet {
		logger.Errorf("Method (%v) is not allowed", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	logger.Info("Redirecting to original url")

	shortUrl := r.URL.Path
	logger.Infof("Short URL (%v)", shortUrl)

	originalUrl, err := h.get.Get(ctx, shortUrl)
	if err != nil {
		logger.Errorf("Error getting original url (%v)", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	logger.Infof("Original URL (%v)", originalUrl)
	w.WriteHeader(http.StatusFound)
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
