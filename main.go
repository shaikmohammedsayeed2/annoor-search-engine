package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "annoor-search-engine/docs" // Replace with your project path

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"annoor-search-engine/handlers"
	"annoor-search-engine/models"
)

// @title Library API
// @version 1.0
// @description This is a sample server for managing books, authors, and publishers.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

var db *gorm.DB

func main() {

	dsn := "postgresql://afse_owner:eaNcZ0MiD3bV@ep-round-brook-a5xr80t9.us-east-2.aws.neon.tech/afse"
	// Initialize the database
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Book{}, &models.Author{}, &models.Publisher{})

	// Initialize Gin router
	r := gin.Default()

	// Book routes
	r.POST("/books", func(c *gin.Context) { handlers.CreateBook(c, db) })
	r.GET("/books", func(c *gin.Context) { handlers.GetBooks(c, db) })
	r.GET("/books/:id", func(c *gin.Context) { handlers.GetBook(c, db) })
	r.PUT("/books/:id", func(c *gin.Context) { handlers.UpdateBook(c, db) })
	r.DELETE("/books/:id", func(c *gin.Context) { handlers.DeleteBook(c, db) })

	// Author routes
	r.POST("/authors", func(c *gin.Context) { handlers.CreateAuthor(c, db) })
	r.GET("/authors", func(c *gin.Context) { handlers.GetAuthors(c, db) })
	r.GET("/authors/:id", func(c *gin.Context) { handlers.GetAuthor(c, db) })
	r.PUT("/authors/:id", func(c *gin.Context) { handlers.UpdateAuthor(c, db) })
	r.DELETE("/authors/:id", func(c *gin.Context) { handlers.DeleteAuthor(c, db) })

	// Publisher routes
	r.POST("/publishers", func(c *gin.Context) { handlers.CreatePublisher(c, db) })
	r.GET("/publishers", func(c *gin.Context) { handlers.GetPublishers(c, db) })
	r.GET("/publishers/:id", func(c *gin.Context) { handlers.GetPublisher(c, db) })
	r.PUT("/publishers/:id", func(c *gin.Context) { handlers.UpdatePublisher(c, db) })
	r.DELETE("/publishers/:id", func(c *gin.Context) { handlers.DeletePublisher(c, db) })

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	r.Run(":8080")
}
