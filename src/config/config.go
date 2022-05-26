package infrastructure

import (
  "os"
  "golang.org/x/oauth2"
)

func LoadConfig() (conf oauth2.Config) {
  conf = oauth2.Config{
    ClientID:     os.Getenv("CLIENT_ID"),
    ClientSecret: os.Getenv("CLIENT_SECRET"),
    Scopes:       []string{"tweet.read", "tweet.write", "users.read"},
    Endpoint: oauth2.Endpoint{
      AuthURL:  "https://twitter.com/i/oauth2/authorize",
      TokenURL: "https://api.twitter.com/2/oauth2/token",
      AuthStyle: oauth2.AuthStyleInHeader,
    },
    RedirectURL: os.Getenv("REDIRECT_URL"),
  }
  return
}
