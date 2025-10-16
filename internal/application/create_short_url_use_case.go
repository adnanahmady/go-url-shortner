package application

import (
	"context"
	"errors"
	"fmt"

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

	shortCode, err := u.repository.Create(ctx, url)
	if err != nil {
		if errors.Is(err, domain.ErrUrlAlreadyExists) {
			lgr.Infof("Short url already exists (%v)", url)
			return toShortUrl(shortCode), ErrAlreadyShorten
		}

		lgr.Errorf("Error creating short url (%v)", err)
		return "", err
	}

	lgr.Infof("Short url created (%v)", shortCode)
	return toShortUrl(shortCode), nil
}

func toShortUrl(shortCode string) string {
	return fmt.Sprintf("http://localhost:5000/%v", shortCode)
}
