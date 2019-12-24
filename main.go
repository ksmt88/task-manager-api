package main

import (
	"./internal/task/infrastructure"
)

func main() {
	infrastructure.Router.Logger.Fatal(
		infrastructure.Router.Start(":1323"))
}
