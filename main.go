package main

import (
	_ "github.com/lib/pq"
	"os"
)

func main() {

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8000")

}
