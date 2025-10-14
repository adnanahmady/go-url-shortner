package domain

import "context"

type UrlRepository interface {
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, shortUrl string) (string, error)
}