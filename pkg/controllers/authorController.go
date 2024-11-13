package controllers

import (
	"github/mohanapranes/book_trust_go/pkg/entities"
	"github/mohanapranes/book_trust_go/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	AuthorService *services.AuthorService
}

func NewAuthorController(AuthorService *services.AuthorService) *AuthorController {
	return &AuthorController{AuthorService: AuthorService}
}

func (h *AuthorController) CreateAuthor(c *gin.Context) {
	var Author entities.Author
	if err := c.ShouldBindJSON(&Author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthorService.CreateAuthor(&Author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create a Author"})
		return
	}
	c.JSON(http.StatusCreated, Author)
}

func (h *AuthorController) GetAllAuthors(c *gin.Context) {
	Authors, err := h.AuthorService.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Authors"})
		return
	}
	c.JSON(http.StatusOK, Authors)
}

func (h *AuthorController) GetAuthorByID(c *gin.Context) {
	AuthorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	Author, err := h.AuthorService.GetAuthorByID(AuthorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, Author)
}

func (h *AuthorController) UpdateAuthor(c *gin.Context) {
	AuthorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	var Author entities.Author
	if err := c.ShouldBindJSON(&Author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthorService.UpdateAuthorByID(AuthorID, &Author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Author"})
		return
	}
	c.JSON(http.StatusOK, Author)
}

func (h *AuthorController) DeleteAuthor(c *gin.Context) {
	AuthorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	if err := h.AuthorService.DeleteAuthor(AuthorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Author"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
