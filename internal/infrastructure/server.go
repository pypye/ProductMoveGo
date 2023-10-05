package infrastructure

import (
	"github.com/gin-gonic/gin"
	"product_move/internal/middleware"
)

type Server struct {
	r *gin.Engine
}

func NewServer() *Server {
	return &Server{r: gin.Default()}
}

func GetMapping(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.GET(path, middleware.Authenticate, handler)
}

func PostMapping(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.POST(path, middleware.Authenticate, handler)
}

func PutMapping(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.PUT(path, middleware.Authenticate, handler)
}

func DeleteMapping(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.DELETE(path, middleware.Authenticate, handler)
}

func GetMappingNoAuth(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.GET(path, handler)
}

func PostMappingNoAuth(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.POST(path, handler)
}

func PutMappingNoAuth(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.PUT(path, handler)
}

func DeleteMappingNoAuth(path string, handler func(ctx *gin.Context)) {
	instance.Server.r.DELETE(path, handler)
}

func (s *Server) Get() *gin.Engine {
	return s.r
}

func (s *Server) Run(port string) error {
	return s.r.Run(port)
}
