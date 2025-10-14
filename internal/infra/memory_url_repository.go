package infra

import (
	"context"

	"github.com/adnanahmady/go-url-shortner/internal/domain"
	request "github.com/adnanahmady/go-url-shortner/pkg/reqeust"
	"github.com/adnanahmady/go-url-shortner/pkg/store"
)

var _ domain.UrlRepository = &MemoryUrlRepository{}

type MemoryUrlRepository struct {
	storeMgr store.StoreManager
}

func NewMemoryUrlRepository(s store.StoreManager) *MemoryUrlRepository {
	return &MemoryUrlRepository{storeMgr: s}
}

func (r *MemoryUrlRepository) Create(ctx context.Context, url string) (string, error) {
	lgr := request.GetLogger(ctx).Section("MemoryUrlRepository", "Create")
	lgr.Infof("Creating short url (%v)", url)
	// TODO create short url for the long url
	err := r.storeMgr.Set(url, url)
	if err != nil {
		lgr.Errorf("Error creating short url (%v)", err)
		return "", err
	}
	lgr.Infof("Short url created (%v)", url)
	return url, nil
}

func (r *MemoryUrlRepository) Get(ctx context.Context, shortUrl string) (string, error) {
	lgr := request.GetLogger(ctx).Section("MemoryUrlRepository", "Get")
	lgr.Infof("Getting short url (%v)", shortUrl)
	url, err := r.storeMgr.Get(shortUrl)
	if err != nil {
		lgr.Errorf("Error getting short url (%v)", err)
		return "", err
	}
	lgr.Infof("Short url got (%v)", url)
	return url, nil
}