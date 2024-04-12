package main

import (
	"fmt"
	"strings"

	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/user"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	walletHandler := wallet.New(p)
	walletGroup := e.Group("/api/v1/wallets")
	walletGroup.GET("", walletHandler.GetWallet)
	walletGroup.POST("", walletHandler.CreateWallet)
	walletGroup.PUT("/:id", walletHandler.UpdateWallet)

	userHandler := user.New(p)
	userGroup := e.Group("/api/v1/users")
	userGroup.GET("/:id/wallets", userHandler.WalletByUserId)

	e.Logger.Fatal(e.Start(":1323"))
}
