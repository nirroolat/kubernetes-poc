package main

import (
	"os"

	"scm.bluebeam.com/stu/golang-template/app"
)

func main() {
	a := app.App{}
	a.Initialize(os.Getenv("DSN"))
	a.Run(":8080")
}
