package application

import (
	"context"

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

func (u *GetShortUrlUseCaseImpl) Get(ctx context.Context, shortUrl string) (string, error) {
	lgr := request.GetLogger(ctx).Section("GetShortUrlUseCaseImpl", "Get")
	lgr.Infof("Getting short url (%v)", shortUrl)
	return u.repository.Get(ctx, shortUrl)
}