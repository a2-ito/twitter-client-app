package usecases

import (
	"context"
	"fmt"
	"io"

	//"os"
	"encoding/json"

	"golang.org/x/oauth2"
)

type Me struct {
	Data User
}

type User struct {
	Id       string
	Name     string
	Username string
}

func (u *appUseCase) Callback(ctx context.Context, code string, queryState string) (result Me, url string) {

	fmt.Println("usecase Callback")
	if code == "" {
		//log.Println("code not found")
		//w.WriteHeader(http.StatusBadRequest)
		return
	}
	if state == "" {
		//log.Println("state not found")
		//w.WriteHeader(http.StatusBadRequest)
		return
	}
	if queryState != state {
		//log.Println("invalid state")
		//w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := conf.Exchange(
		ctx,
		code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}
	token = t
	fmt.Println("token: ", token)

	oAuthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	res, err := oAuthClient.Get("https://api.twitter.com/2/users/me")
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
	fmt.Println(string(body))
	var me Me
	if err := json.Unmarshal(body, &me); err != nil {
		fmt.Printf("failed to get me: %v\n", err)
	}
	fmt.Println("content: ", me)

	// Todo: 定数化
	url = "http://localhost:8080/?id=" + me.Data.Id + "&username=" + me.Data.Username

	return me, url

}
