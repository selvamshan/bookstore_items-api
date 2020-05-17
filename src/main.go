package main

import (
	"os"
	"github.com/selvamshan/bookstore_items-api/src/app"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApplication()
}
