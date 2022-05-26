package usecases

import (
  "fmt"
  "context"
  "io"
  "os"
  "golang.org/x/oauth2"
)

func (u *appUseCase) Callback(ctx context.Context, code string) (err error) {

  fmt.Println("usecase Callback")
  token, err := conf.Exchange(
                  ctx,
                  code,
                  oauth2.SetAuthURLParam("code_verifier", codeVerifier))
  fmt.Println("token: ", token)
  if err != nil {
    fmt.Println(err)
    //log.Fatal(err)
  }

  oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
  res, err := oAuthClient.Get("https://api.twitter.com/2/users/me")
  if err != nil {
    //log.Printf("failed to get me: %v\n", err)
    fmt.Printf("failed to get me: %v\n", err)
    //w.WriteHeader(http.StatusInternalServerError)
  }

  defer res.Body.Close()
  io.Copy(os.Stdout, res.Body)

  b, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("failed to get me: %v\n", err)
  }
  fmt.Println("content: ", string(b))

  return nil

}
