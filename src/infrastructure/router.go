package infrastructure

import (
  "net/http"
  "io"
  "fmt"
  "crypto/rand"
  //"crypto/sha256"
  "encoding/base64"
  //"context"

  "github.com/joho/godotenv"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  //"golang.org/x/oauth2"
  handler "src/interfaces"
)

/*
func hello(c echo.Context) error {
  fmt.Println("id ", c.QueryParam("id"))
  return c.String(http.StatusOK, "Hello, World!")
}
*/

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
    e := echo.New()

    // Env
    // .env ファイルが存在する場合はそこから読み込み、存在する場合は、環境変数を参照する
    err := godotenv.Load(".env")
    if err != nil {
      fmt.Println(err)
    }

    handler := handler.NewAppHandler()
    handler.NewAppUseCase()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", handler.Hello)
    e.GET("/login", handler.Login)
    e.GET("/callback", handler.Callback)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
