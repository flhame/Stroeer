package main

import (
	"Strooer/internal/app/userservice"
)
func main() {
	app := userservice.Application{}
	app.Start()
}