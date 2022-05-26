package infrastructure

import (
  "os"
  "fmt"
  "golang.org/x/oauth2"
)

func LoadConfig() (conf oauth2.Config) {

  _, result := os.LookupEnv("CLIENT_ID")
  if (result == false) {
    fmt.Println("Not found: CLIENT_ID")
    // Todo: 5xx を返却
  }
  _, result = os.LookupEnv("CLIENT_SECRET")
  if (result == false) {
    fmt.Println("Not found: CLIENT_SECRET")
    // Todo: 5xx を返却
  }
  _, result = os.LookupEnv("REDIRECT_URL")
  if (result == false) {
    fmt.Println("Not found: REDIRECT_URL")
    // Todo: 5xx を返却
  }

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