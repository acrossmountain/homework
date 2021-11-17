package controller

import (
	"net/http"
	"strings"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
)

func init() {
	gs.Provide(new(Index)).Init(func(c *Index) {
		gs.GetMapping("/header", c.HeaderHandle)
		gs.GetMapping("/health", c.HealthHandle)
		gs.GetMapping("/version", c.VersionHandle)
	})
}

type Index struct {
	Version string `value:"${VERSION}"`
}

func (c *Index) HeaderHandle(ctx web.Context) {

	for key, value := range ctx.Request().Header {
		ctx.Header(key, strings.Join(value, ""))
	}
	ctx.Status(http.StatusOK)
	ctx.String("ok")
}

func (c *Index) HealthHandle(ctx web.Context) {

	ctx.Status(http.StatusOK)
	ctx.String("ok")
}

func (c *Index) VersionHandle(ctx web.Context) {

	ctx.Status(http.StatusOK)
	ctx.String(c.Version)
}
