package usecases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/oauth2"
)

type Like struct {
	Tweet_id string `json:"tweet_id"`
}

func (u *appUseCase) Likes(ctx context.Context, id string, tweet_id string) error {
	fmt.Println("useCase Likes")
	url := "https://api.twitter.com/2/users/" + id + "/likes"

	oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

	l := new(Like)
	l.Tweet_id = tweet_id
	like, _ := json.Marshal(l)

	res, err := oAuthClient.Post(url, "application/json", bytes.NewBuffer(like))
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
