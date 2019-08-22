package apis

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

var echod *echo.Echo

func init() {
	echod = echo.New()
	echod.GET("hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
}

// Run api server
func Run() {
	echod.Logger.Fatal(echod.Start(viper.GetString("bind")))
}

// APIs return echo pointer
func APIs() *echo.Echo {
	return echod
}
