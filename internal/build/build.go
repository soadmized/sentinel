package build

import (
	"github.com/labstack/echo/v4"
	"sentry/internal/api"
	"sentry/internal/config"
	"sentry/internal/repo"
	"sentry/internal/service"
)

type Builder struct {
	conf config.Config
}

func New(conf config.Config) (*Builder, error) {
	b := Builder{conf: conf}

	return &b, nil
}

func (b *Builder) Api() (*api.Api, error) {
	srv, err := b.service()
	if err != nil {
		return nil, err
	}

	server := echo.New()
	a := api.Api{
		Service: srv,
		Server:  server,
	}

	return &a, nil
}

func (b *Builder) service() (*service.Service, error) {
	r, err := b.repo()
	if err != nil {
		return nil, err
	}

	srv := service.Service{Repo: r}

	return &srv, nil
}

func (b *Builder) repo() (*repo.Repo, error) {
	r, err := repo.New(b.conf)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
