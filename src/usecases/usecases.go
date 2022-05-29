package usecases

import (
  "context"
  //"github.com/labstack/echo/v4"
  "golang.org/x/oauth2"
)

type AppUseCase interface {
  Hello(ctx context.Context) error
  Login(ctx context.Context) string
  Callback(ctx context.Context, code string, queryState string) (result Me, url string)
  Timelines(ctx context.Context, id string) (result Timelines)
  Tweet(ctx context.Context, tweetTest string) error
  Follow(ctx context.Context, id string, target_id string) error
  Search(ctx context.Context, query string) error
  Likes(ctx context.Context, id string, tweet_id string) error
}

type appUseCase struct {
}

var conf oauth2.Config
var codeVerifier string
var token *oauth2.Token
var state string
var tokenStore map[string]string

func NewAppUseCase() AppUseCase {
  return &appUseCase{}
}

