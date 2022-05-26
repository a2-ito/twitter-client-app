package infrastructure

import (
  "net/http"
  "io"
  "fmt"
  "crypto/rand"
  "crypto/sha256"
  "encoding/base64"
  "context"

  "github.com/joho/godotenv"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "golang.org/x/oauth2"
)

func hello(c echo.Context) error {
  fmt.Println("id ", c.QueryParam("id"))
  return c.String(http.StatusOK, "Hello, World!")
}

func auth(c echo.Context) error {
  url := "http://localhost:1323/auth/hoge"
  return c.Redirect(http.StatusMovedPermanently, url)
}

func callbackHandler(c echo.Context) error {
  fmt.Println("state: ", c.QueryParam("state"))
  fmt.Println("code:  ", c.QueryParam("code"))
  return c.String(http.StatusOK, "callback!")
}

func generateBase64Encoded32byteRandomString() string {
  b := make([]byte, 32)
  if _, err := io.ReadFull(rand.Reader, b); err != nil {
    return ""
  }
  return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}

func Run() {
    ctx := context.Background()
    // Echo instance
    e := echo.New()
    /*
    i := interactor.NewInteractor()
    h := i.NewUserHandler()
    */

    // Env
    err := godotenv.Load(".env")
    if err != nil {
      fmt.Println(err)
    }

    // Config
    conf := LoadConfig()

    codeVerifier := generateBase64Encoded32byteRandomString()
    //codeVerifier := "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"
    //codeChallenge := "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"
    h := sha256.New()
    h.Write([]byte(codeVerifier))
    hashed := h.Sum(nil)
    codeChallenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hashed)

    authURL := conf.AuthCodeURL(
                "FIXME",
                oauth2.SetAuthURLParam("code_challenge", codeChallenge),
                oauth2.SetAuthURLParam("code_challenge_method", "S256"))

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", hello)
    //e.GET("/auth", auth)
    e.GET("/login", func(c echo.Context) error {
      return c.Redirect(http.StatusMovedPermanently, authURL)
    })

    //e.GET("/callback", callbackHandler)
    e.GET("/callback", func(c echo.Context) error {
      code := c.QueryParam("code")
      fmt.Println("code: ", c.QueryParam("code"))
      fmt.Println("codeVerifier: ", codeVerifier)
      token, err := conf.Exchange(
                      ctx,
                      code,
                      oauth2.SetAuthURLParam("code_verifier", codeVerifier))
      if err != nil {
        fmt.Println(err)
        //log.Fatal(err)
      }
      fmt.Println("token: ", token)
      return c.String(http.StatusOK, "callback!")
    })

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
