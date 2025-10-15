package application

import (
	"context"
	"errors"
	"strings"

	"github.com/adnanahmady/go-url-shortner/internal/domain"
	request "github.com/adnanahmady/go-url-shortner/pkg/reqeust"
)

type GetShortUrlUseCase interface {
	Get(ctx context.Context, shortUrl string) (string, error)
}

type GetShortUrlUseCaseImpl struct {
	repository domain.UrlRepository
}

func NewGetShortUrlUseCaseImpl(repository domain.UrlRepository) *GetShortUrlUseCaseImpl {
	return &GetShortUrlUseCaseImpl{repository: repository}
}

func (u *GetShortUrlUseCaseImpl) Get(ctx context.Context, shortParam string) (string, error) {
	lgr := request.GetLogger(ctx).Section("GetShortUrlUseCaseImpl", "Get")
	lgr.Infof("Getting short url (%v)", shortParam)

	shortCode := strings.TrimPrefix(shortParam, "/")
	url, err := u.repository.Get(ctx, shortCode)
	if err != nil {
		if errors.Is(err, domain.ErrUrlNotFound) {
			lgr.Infof("Short url not found (%v)", shortCode)
			return "", ErrUrlNotFound
		}
		lgr.Errorf("Error getting short url (%v)", err)
		return "", err
	}
	return url, nil
}