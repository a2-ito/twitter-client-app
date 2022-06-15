package infrastructure

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	//"golang.org/x/oauth2"
	handler "interfaces"
)

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

	// Todo: 暫定対応 ドメイン取得後CORS許可は解除すること
	e.Use(middleware.CORS())

	e.GET("/", handler.Hello)
	e.GET("/login", handler.Login)
	e.GET("/callback", handler.Callback)
	e.GET("/timelines", handler.Timelines)
	e.POST("/tweet", handler.Tweet)
	e.POST("/follow", handler.Follow)
	e.GET("/search", handler.Search)
	e.POST("/likes", handler.Likes)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
