package usecases

import (
  "fmt"
  "context"
  "io"
  "os"
  "bytes"
  "encoding/json"
  "golang.org/x/oauth2"
)

type Post struct {
  Text string `json:"text"`
}

func (u *appUseCase) Tweet(ctx context.Context, tweetText string) error {
  fmt.Println("useCase Tweet")
  url := "https://api.twitter.com/2/tweets"

  fmt.Println("token: ", token)

  oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

  post := new(Post)
  post.Text = tweetText
  tweet, _ := json.Marshal(post)

  res, err := oAuthClient.Post(url, "application/json", bytes.NewBuffer(tweet))
  if err != nil {
    //log.Printf("failed to get me: %v\n", err)
    fmt.Printf("failed to get me: %v\n", err)
    //w.WriteHeader(http.StatusInternalServerError)
  }
  defer res.Body.Close()

  io.Copy(os.Stdout, res.Body)
  /*
  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("failed to get me: %v\n", err)
  }

  var tl Timelines
  if err := json.Unmarshal(body, &tl); err != nil {
    fmt.Printf("failed to get me: %v\n", err)
  }

  fmt.Printf("%-v \n", tl)
  */

  return nil
}
