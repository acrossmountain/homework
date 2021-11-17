package chain

import (
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
)

func init() {
	gs.Object(&LoggerFilter{}).Export((*web.Filter)(nil))
}

type LoggerFilter struct {
}

func (c *LoggerFilter) Invoke(ctx web.Context, chain web.FilterChain) {
	chain.Next(ctx)
	w := ctx.ResponseWriter()
	log.Infof("remote ip: %s, response status code: %d", ctx.ClientIP(), w.Status())
}
