package main

import (
	_ "github.com/acrossmountain/homework/module08/internal/controller"
	_ "github.com/go-spring/starter-echo"

	"github.com/go-spring/spring-core/gs"
)

func main() {
	gs.Property("VERSION", "v1.1.0-rc2")
	gs.Run()
}
