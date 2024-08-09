package handlers

import (
	"net/http"

	"annoor-search-engine/models" // Replace with your actual project path

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Create an author
// @Description Create a new author
// @Tags authors
// @Accept  json
// @Produce  json
// @Param author body models.Author true "Author"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.ErrorResponse
// @Router /authors [post]
func CreateAuthor(c *gin.Context, db *gorm.DB) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db.Create(&author)
	c.JSON(http.StatusOK, author)
}

// @Summary Get all authors
// @Description Get all authors
// @Tags authors
// @Produce  json
// @Success 200 {array} models.Author
// @Router /authors [get]
func GetAuthors(c *gin.Context, db *gorm.DB) {
	var authors []models.Author
	db.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

// @Summary Get an author by ID
// @Description Get a single author by ID
// @Tags authors
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Failure 404 {object} models.ErrorResponse
// @Router /authors/{id} [get]
func GetAuthor(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

// @Summary Update an author by ID
// @Description Update an author by their ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Param author body models.Author true "Author"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /authors/{id} [put]
func UpdateAuthor(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Author not found"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	db.Save(&author)
	c.JSON(http.StatusOK, author)
}

// @Summary Delete an author by ID
// @Description Delete an author by their ID
// @Tags authors
// @Param id path int true "Author ID"
// @Success 204
// @Failure 404 {object} models.ErrorResponse
// @Router /authors/{id} [delete]
func DeleteAuthor(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var author models.Author
	if err := db.Delete(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Author not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
