package usecases

import (
  "fmt"
  "context"
  "golang.org/x/oauth2"
)

func (u *appUseCase) Callback(ctx context.Context, code string) (err error) {

  token, err := conf.Exchange(
                  ctx,
                  code,
                  oauth2.SetAuthURLParam("code_verifier", codeVerifier))
  fmt.Println(token)
  if err != nil {
    fmt.Println(err)
    //log.Fatal(err)
  }

  return nil

}
