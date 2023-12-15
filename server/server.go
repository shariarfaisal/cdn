package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	uploadDir string
	address   string
	router    *gin.Engine
	cache     *CachingStore
}

func NewServer(address string) *Server {
	server := &Server{
		uploadDir: "./bucket",
		address:   address,
		router:    gin.Default(),
		cache:     NewCachingStore(time.Minute * 10),
	}

	server.setupRouter()

	return server
}

func (server *Server) Start() error {
	fmt.Printf("Server is running on http://localhost%s\n", server.address)
	return server.router.Run(server.address)
}

func (s *Server) setupRouter() {
	config := cors.DefaultConfig()

	// Allow all origins
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	s.router.Use(cors.New(config))

	r := s.router.Group("/")

	// Set up routes
	r.POST("/", s.uploadFile)
	r.GET("/", s.getFiles)
	r.GET("/:filename", s.getFile)
	r.DELETE("/:filename", s.deleteFile)

	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
}
