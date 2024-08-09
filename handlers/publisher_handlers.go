package handlers

import (
	"net/http"

	"annoor-search-engine/models" // Replace with your actual project path

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Create a publisher
// @Description Create a new publisher
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param publisher body models.Publisher true "Publisher"
// @Success 200 {object} models.Publisher
// @Failure 400 {object} models.ErrorResponse
// @Router /publishers [post]
func CreatePublisher(c *gin.Context, db *gorm.DB) {
	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	db.Create(&publisher)
	c.JSON(http.StatusOK, publisher)
}

// @Summary Get all publishers
// @Description Get all publishers
// @Tags publishers
// @Produce  json
// @Success 200 {array} models.Publisher
// @Router /publishers [get]
func GetPublishers(c *gin.Context, db *gorm.DB) {
	var publishers []models.Publisher
	db.Find(&publishers)
	c.JSON(http.StatusOK, publishers)
}

// @Summary Get a publisher by ID
// @Description Get a single publisher by ID
// @Tags publishers
// @Produce  json
// @Param id path int true "Publisher ID"
// @Success 200 {object} models.Publisher
// @Failure 404 {object} models.ErrorResponse
// @Router /publishers/{id} [get]
func GetPublisher(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var publisher models.Publisher
	if err := db.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Publisher not found"})
		return
	}
	c.JSON(http.StatusOK, publisher)
}

// @Summary Update a publisher by ID
// @Description Update a publisher by its ID
// @Tags publishers
// @Accept  json
// @Produce  json
// @Param id path int true "Publisher ID"
// @Param publisher body models.Publisher true "Publisher"
// @Success 200 {object} models.Publisher
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /publishers/{id} [put]
func UpdatePublisher(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var publisher models.Publisher
	if err := db.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Publisher not found"})
		return
	}

	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	db.Save(&publisher)
	c.JSON(http.StatusOK, publisher)
}

// @Summary Delete a publisher by ID
// @Description Delete a publisher by its ID
// @Tags publishers
// @Param id path int true "Publisher ID"
// @Success 204
// @Failure 404 {object} models.ErrorResponse
// @Router /publishers/{id} [delete]
func DeletePublisher(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var publisher models.Publisher
	if err := db.Delete(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Publisher not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
