package usecases

import (
	"context"
	"fmt"
	"io"

	//"os"
	"encoding/json"

	"golang.org/x/oauth2"
)

type Timelines struct {
	Data []Tweet `json:"data"`
	Meta Meta    `json:"meta"`
}

type Meta struct {
	Result_count uint64 `json:"result_count"`
	Newest_id    string `json:"newest_id"`
	Oldest_id    string `json:"oldest_id"`
	Next_token   string `json:"next_token"`
}

type Tweet struct {
	Id         string `json:"id"`
	Text       string `json:"text"`
	Created_at string `json:"created_at"`
	Author_id  string `json:"author_id"`
}

func (u *appUseCase) Timelines(ctx context.Context, id string) (result Timelines) {
	fmt.Println("useCase Timelines")
	url := "https://api.twitter.com/2/users/" + id + "/timelines/reverse_chronological?expansions=author_id&tweet.fields=conversation_id,created_at&user.fields=entities&max_results=10"
	//url := "https://api.twitter.com/2/users/"+id+"/tweets"

	fmt.Println("token: ", token)

	oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	res, err := oAuthClient.Get(url)
	if err != nil {
		//log.Printf("failed to get me: %v\n", err)
		fmt.Printf("failed to get me: %v\n", err)
		//w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	//io.Copy(os.Stdout, res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("failed to get me: %v\n", err)
	}

	var tl Timelines
	if err := json.Unmarshal(body, &tl); err != nil {
		fmt.Printf("failed to get me: %v\n", err)
	}

	fmt.Printf("%-v \n", tl)

	return tl
}
