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

type Post struct {
	Text string `json:"text"`
}

func (u *appUseCase) Tweet(ctx context.Context, tweetText string) error {
	fmt.Println("useCase Tweet")

	// Todo: URL情報の定数化
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

        // Todo: Tweet後レスポンスコードをクライアントへ返却
	io.Copy(os.Stdout, res.Body)

	return nil
}
