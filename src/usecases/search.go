package usecases

import (
  "fmt"
  "context"
  "io"
  //"os"
  "encoding/json"
  "golang.org/x/oauth2"
)

type Results struct {
  Data []Tweet `json:"data"`
}

// Notes:
// Meta, Tweet は timelines で宣言済み

func (u *appUseCase) Search(ctx context.Context, query string) (result Results) {
  fmt.Println("useCase Search")
  url := "https://api.twitter.com/2/tweets/search/recent?query="+query+"&expansions=author_id&tweet.fields=conversation_id,created_at&user.fields=entities&&max_results=10"

  fmt.Println("token: ", token)

  oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

  res, err := oAuthClient.Get(url)
  if err != nil {
    //log.Printf("failed to get me: %v\n", err)
    fmt.Printf("failed to get me: %v\n", err)
    //w.WriteHeader(http.StatusInternalServerError)
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("failed to get me: %v\n", err)
  }

  var r Results
  if err := json.Unmarshal(body, &r); err != nil {
    fmt.Printf("failed to get me: %v\n", err)
  }

  fmt.Printf("%-v \n", r)

  return r
}
