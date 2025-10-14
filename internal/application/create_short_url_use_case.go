package application

import (
	"context"

	"github.com/adnanahmady/go-url-shortner/internal/domain"
	request "github.com/adnanahmady/go-url-shortner/pkg/reqeust"
)

type CreateShortUrlUseCase interface {
	Create(ctx context.Context, url string) (string, error)
}

var _ CreateShortUrlUseCase = (*CreateShortUrlUseCaseImpl)(nil)

type CreateShortUrlUseCaseImpl struct {
	repository domain.UrlRepository
}

func NewCreateShortUrlUseCaseImpl(repository domain.UrlRepository) *CreateShortUrlUseCaseImpl {
	return &CreateShortUrlUseCaseImpl{repository: repository}
}

func (u *CreateShortUrlUseCaseImpl) Create(ctx context.Context, url string) (string, error) {
	lgr := request.GetLogger(ctx).Section("CreateShortUrlUseCaseImpl", "Create")
	lgr.Infof("Creating short url (%v)", url)

	return u.repository.Create(ctx, url)
}