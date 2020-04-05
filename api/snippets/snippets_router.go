package snippets

import (
	"github.com/gin-gonic/gin"
)

//Routes holds all user routes
func Routes(route *gin.Engine) {
	router := route.Group("/snippets")
	router.POST("/", CreateSnippet)
	router.GET("/:id", GetSnippetByID)
	// router.GET("/:userId", GetUserSnippets)
	// router.PUT("/:id", UpdateSnippet)
	// router.DELETE("/:id", DeleteSnippet)

}
