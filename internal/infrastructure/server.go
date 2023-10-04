package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	r *gin.Engine
}

type ServerCtx struct {
	ctx *gin.Context
}

func NewServer() *Server {
	return &Server{r: gin.Default()}
}

func GetMapping(path string, handler func(ctx *ServerCtx)) {
	instance.server.r.GET(path, func(c *gin.Context) {
		handler(&ServerCtx{ctx: c})
	})
}

func PostMapping(path string, handler func(ctx *ServerCtx)) {
	instance.server.r.POST(path, func(c *gin.Context) {
		handler(&ServerCtx{ctx: c})
	})
}

func PutMapping(path string, handler func(ctx *ServerCtx)) {
	instance.server.r.PUT(path, func(c *gin.Context) {
		handler(&ServerCtx{ctx: c})
	})
}

func DeleteMapping(path string, handler func(ctx *ServerCtx)) {
	instance.server.r.DELETE(path, func(c *gin.Context) {
		handler(&ServerCtx{ctx: c})
	})
}

func (s *Server) Run(port string) error {
	return s.r.Run(port)
}

func (c *ServerCtx) Param(key string) string {
	return c.ctx.Param(key)
}

func (c *ServerCtx) JSON(code int, obj interface{}) {
	c.ctx.JSON(code, gin.H{
		"code":    code,
		"message": http.StatusText(code),
		"result":  obj,
	})
}
