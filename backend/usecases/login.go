package usecases

import (
	config "config"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"golang.org/x/oauth2"
)

func generateBase64Encoded32byteRandomString() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}

func getAuthURL() string {

	codeVerifier = generateBase64Encoded32byteRandomString()
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	hashed := h.Sum(nil)
	codeChallenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hashed)

	// init state
	state = fmt.Sprintf("st%d", time.Now().UnixNano())

	authURL := conf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"))
	return authURL
}

func (u *appUseCase) Login(ctx context.Context) string {

	fmt.Println("usecase Login")
	conf = config.LoadConfig()
	url := getAuthURL()
	fmt.Println("url: ", url)
	return url

}
