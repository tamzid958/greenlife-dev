package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {

	a := App{}

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	a.Initialize(
		fmt.Sprint(viper.Get("APP_DB_USERNAME")),
		fmt.Sprint(viper.Get("APP_DB_PASSWORD")),
		fmt.Sprint(viper.Get("APP_DB_NAME")),
	)

	a.Run(":8000")

}
