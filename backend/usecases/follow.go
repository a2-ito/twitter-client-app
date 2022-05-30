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

type FollowingInfo struct {
	Target_user_id string `json:"target_user_id"`
}

func (u *appUseCase) Follow(ctx context.Context, id string, target_id string) error {
	fmt.Println("useCase Follow ", id, target_id)

	// Todo: URL情報の定数化
	url := "https://api.twitter.com/2/users/" + id + "/following"

	oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

	f := new(FollowingInfo)
	f.Target_user_id = target_id
	following, _ := json.Marshal(f)

	res, err := oAuthClient.Post(url, "application/json", bytes.NewBuffer(following))
	if err != nil {
		//log.Printf("failed to get me: %v\n", err)
		fmt.Printf("failed to get me: %v\n", err)
		//w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

        // Todo: レスポンスコードをクライアントへ返却
	io.Copy(os.Stdout, res.Body)

	return nil
}
