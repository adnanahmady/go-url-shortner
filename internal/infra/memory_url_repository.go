package infra

import (
	"context"
	"errors"
	"math/rand"
	"time"

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
	r.storeMgr.Lock()
	defer r.storeMgr.Unlock()

	if shortCode, ok := r.storeMgr.HasUrl(url); ok {
		lgr.Infof("Short url already exists (%v)", url)
		return shortCode, domain.ErrUrlAlreadyExists
	}

	shortCode := r.generateShortCode()
	err := r.storeMgr.Set(shortCode, url)
	if err != nil {
		lgr.Errorf("Error creating short url (%v)", err)
		return "", err
	}
	lgr.Infof("Short url created (%v)", url)
	return shortCode, nil
}

func (r *MemoryUrlRepository) Get(ctx context.Context, shortUrl string) (string, error) {
	lgr := request.GetLogger(ctx).Section("MemoryUrlRepository", "Get")
	lgr.Infof("Getting short url (%v)", shortUrl)

	r.storeMgr.Lock()
	defer r.storeMgr.Unlock()

	url, err := r.storeMgr.Get(shortUrl)
	if err != nil {
		if errors.Is(err, store.ErrKeyNotFound) {
			lgr.Infof("Short url not found (%v)", shortUrl)
			return "", domain.ErrUrlNotFound
		}
		lgr.Errorf("Error getting short url (%v)", err)
		return "", err
	}
	lgr.Infof("Short url got (%v)", url)
	return url, nil
}

func (r *MemoryUrlRepository) generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = rune(charset[seededRand.Intn(len(charset))])
	}
	return string(runes)
}
