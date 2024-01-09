package build

import (
	"crypto/subtle"
	"sentinel/internal/api"
	"sentinel/internal/config"
	"sentinel/internal/repo"
	"sentinel/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	server.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(b.conf.AppUser)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(b.conf.AppPass)) == 1 {
			return true, nil
		}

		return false, nil
	}))

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
