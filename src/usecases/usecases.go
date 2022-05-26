package usecases

import (
  "context"
  //"github.com/labstack/echo/v4"
  "golang.org/x/oauth2"
)

type AppUseCase interface {
  Hello(ctx context.Context) error
  Login(ctx context.Context) string
  Callback(ctx context.Context, code string) error
}

type appUseCase struct {
}

var conf oauth2.Config
var codeVerifier string

func NewAppUseCase() AppUseCase {
  return &appUseCase{}
}

