package handler

import (
  "context"
  "net/http"
  "fmt"
  "github.com/labstack/echo/v4"
  usecases "src/usecases"
  //"github.com/a2-ito/go-onion/src/domain/model"
)

type AppHandler interface {
  NewAppUseCase() error
  Hello(c echo.Context) error
  Login(c echo.Context) error
  Callback(c echo.Context) error
}

type appHandler struct {
  AppUseCase usecases.AppUseCase
}

func NewAppHandler() AppHandler {
  return &appHandler{}
}

func (h *appHandler) NewAppUseCase() error {
  h.AppUseCase = usecases.NewAppUseCase()
  return nil
}

func (h *appHandler) Hello(c echo.Context) error {
  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  h.AppUseCase.Hello(ctx)
  return c.String(http.StatusOK, "Hello, World!")
}

func (h *appHandler) Login(c echo.Context) error {
  fmt.Println("appHandler Login")

  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  url := h.AppUseCase.Login(ctx)
  return c.Redirect(http.StatusMovedPermanently, url)
}

func (h *appHandler) Callback(c echo.Context) error {
  fmt.Println("appHandler Login")

  code := c.QueryParam("code")
  fmt.Println("code: ", c.QueryParam("code"))

  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  h.AppUseCase.Callback(ctx, code)
  return c.String(http.StatusOK, "callback!")
}

