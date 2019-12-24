package main

import (
	"github.com/ksmt88/taskManager-api/internal/task/infrastructure"
)

func main() {
	infrastructure.Router.Logger.Fatal(
		infrastructure.Router.Start(":1323"))
}
