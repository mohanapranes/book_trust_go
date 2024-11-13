package routes

import (
	"github/mohanapranes/book_trust_go/config/database"
	"github/mohanapranes/book_trust_go/pkg/controllers"
	"github/mohanapranes/book_trust_go/pkg/repository"
	"github/mohanapranes/book_trust_go/pkg/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, database database.Database) {

	authorRepo := repository.NewAuthorRepository(database.GetDb())
	authorService := services.NewAuthorService(authorRepo)
	authorController := controllers.NewAuthorController(authorService)

	routes := router.Group("/api")
	{
		v1Routes := routes.Group("/v1")
		{
			bookRoutes := v1Routes.Group("/books")
			{
				bookRoutes.GET("/")
			}

			authorRoutes := v1Routes.Group("/author")
			{
				authorRoutes.POST("/", authorController.CreateAuthor)
				authorRoutes.GET("/", authorController.GetAllAuthors)
				authorRoutes.GET("/:id", authorController.GetAuthorByID)
				authorRoutes.PUT("/:id", authorController.UpdateAuthor)
				authorRoutes.DELETE("/:id", authorController.DeleteAuthor)
			}
		}
	}

}
