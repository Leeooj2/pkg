package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	GinContextContext = "context" // GinContextContext 存在 gin context 中的标准库 context 实例的 key
	HTTPTraceIDHeader = "x-trace-id"
	ContextTraceID    = contextKey(HTTPTraceIDHeader)
)

func defaultContext(defaultContext context.Context, ctx *gin.Context) (c context.Context) {
	c = defaultContext
	it, b := ctx.Get(GinContextContext)
	if !b {
		log.Warnln("nova-context doesn't exists")
		return
	}
	if c, b = it.(context.Context); !b {
		log.Warnln("invalid nova-context value type")
		c = defaultContext
		return
	}
	return
}

func DefaultTodoContext(ctx *gin.Context) context.Context {
	return defaultContext(context.TODO(), ctx)
}

func ShouldGetTraceID(c context.Context) (traceID string) {
	it := c.Value(ContextTraceID)
	if it == nil {
		// log.Warnln("Could not get trace id from context")
		return
	}
	var ok bool
	if traceID, ok = it.(string); !ok {
		log.Errorf("Invalid trace id value in context: %v", it)
	}
	return
}
