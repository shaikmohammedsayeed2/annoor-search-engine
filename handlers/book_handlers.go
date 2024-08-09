package handlers

import (
	"net/http"

	"annoor-search-engine/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Create a book
// @Description Create a new book
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Router /books [post]
func CreateBook(c *gin.Context, db *gorm.DB) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusOK, book)
}

// @Summary Get all books
// @Description Get all books
// @Tags books
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context, db *gorm.DB) {
	var books []models.Book
	db.Preload("Author").Preload("Publisher").Find(&books)
	c.JSON(http.StatusOK, books)
}

// @Summary Get a book by ID
// @Description Get a single book by ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [get]
func GetBook(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var book models.Book
	if err := db.Preload("Author").Preload("Publisher").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Update a book by ID
// @Description Update a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

// @Summary Delete a book by ID
// @Description Delete a book by its ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 204
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var book models.Book
	if err := db.Delete(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Book not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
