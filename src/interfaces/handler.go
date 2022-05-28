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
  Timelines(c echo.Context) error
  Tweet(c echo.Context) error
}

type appHandler struct {
  AppUseCase usecases.AppUseCase
}

// Tweet 用 struct
type Post struct {
  Text string `json:"text"`
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
  code := c.QueryParam("code")
  fmt.Println("code: ", c.QueryParam("code"))

  queryState := c.QueryParam("state")
  fmt.Println("queryState: ", queryState)

  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  res := h.AppUseCase.Callback(ctx, code, queryState)
  return c.JSON(http.StatusOK, res)
}

func (h *appHandler) Timelines(c echo.Context) error {
  fmt.Println("appHandler Timelines")

  id := c.QueryParam("id")
  fmt.Println("id: ", c.QueryParam("id"))

  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }

  res := h.AppUseCase.Timelines(ctx, id)
  return c.JSON(http.StatusOK, res)
}

func (h *appHandler) Tweet(c echo.Context) error {
  fmt.Println("appHandler Tweet")

  var p Post
  if err := c.Bind(&p); err != nil {
    fmt.Println(err)
  }

  ctx := c.Request().Context()
  if ctx == nil {
    ctx = context.Background()
  }
  h.AppUseCase.Tweet(ctx, p.Text)
  return c.String(http.StatusOK, "tweet!")
}

