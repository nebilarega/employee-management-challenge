package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/graphql/graph"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/graphql/graph/generated"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/middleware"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/models"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/repository"
	"gorm.io/gorm"
)

// Server ...
type Server struct {
	Gin           *gin.Engine
	DB            *gorm.DB
	ModelRegistry *models.Model
}

func NewServer() *Server {
	server := &Server{}
	server.ModelRegistry = models.NewModel()

	if err := server.ModelRegistry.OpenPostgres(); err != nil {
		log.Fatalf("gorm: could not connect to db %q", err)
	}

	server.DB = server.ModelRegistry.DB
	server.ModelRegistry.RegisterAllModels()
	server.ModelRegistry.AutoMigrateAll()

	server.Gin = server.NewRouter()

	return server
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler(server *Server, h *handler.Server) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// NewRouter ...
func (s *Server) NewRouter() *gin.Engine {
	EmployeeRepository := repository.ProvideEmployeeRepository(s.DB)
	DepartmentRepository := repository.ProvideDepartmentRepository(s.DB)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		EmployeeRepository:   EmployeeRepository,
		DepartmentRepository: DepartmentRepository,
	}}))

	r := gin.Default()
	//r.Use(cors.Default())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.GinContextToContextMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Group("/public")
	{

		r.Static("/files", "./files")
	}

	r.GET("/api", playgroundHandler())
	r.POST("/query", graphqlHandler(s, h))

	return r
}

// GetDB returns gorm (ORM)
func (s *Server) GetDB() *gorm.DB {
	return s.DB
}

// Start the http server
func (s *Server) Start() error {
	port := os.Getenv("ADDRESS")

	log.Fatal(s.Gin.Run(":" + port))
	return nil
}

// GracefulShutdown Wait for interrupt signal
// to gracefully shutdown the server with a timeout of 5 seconds.
func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// close database connection
	if s.DB != nil {
		db, _ := s.DB.DB()
		db.Close()
	}
}
