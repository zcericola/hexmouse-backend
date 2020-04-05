package snippets

import (
	"github.com/gin-gonic/gin"
	"github.com/zcericola/hexmouse-backend/api/utils"
	"github.com/zcericola/hexmouse-backend/db"
)

//CreateSnippet adds a new code snippet
func CreateSnippet(c *gin.Context) {
	params := Snippet{}
	c.BindJSON(&params)

	newSnippet := Snippet{}

	params.Snippet = EscapeSpecialCharacters(params.Snippet)

	insertStatement := `
	INSERT INTO snippets
	(snippet, language_id)
	VALUES($1, $2)
	RETURNING snippet_id
	, snippet
	, language_id;`

	stmt, err := db.DB.Prepare(insertStatement)
	utils.HandleError(err)

	err = stmt.QueryRow(
		params.Snippet,
		params.LanguageID,
	).Scan(&newSnippet.SnippetID,
		&newSnippet.Snippet,
		&newSnippet.LanguageID,
	)

	//After creating snippet row, link it here
	//TODO: Get userInfo with middleWare
	// LinkSnippetToUser()

	c.JSON(200, gin.H{
		"message": "Snippet successfully created.",
		"data":    newSnippet,
	})
}

//GetSnippetByID returns a specific snippet
func GetSnippetByID(c *gin.Context) {
	snippetID := c.Param("id")
	snippet := Snippet{}

	selectStatement := `
		SELECT s.snippet
		, s.language_id
		FROM snippets s
		WHERE s.snippet_id = $1
	`
	stmt, err := db.DB.Prepare(selectStatement)
	utils.HandleError(err)

	err = stmt.QueryRow(
		snippetID,
	).Scan(
		&snippet.Snippet,
		&snippet.LanguageID,
	)

	c.JSON(200, gin.H{
		"message": "Snippet successfully retrieved.",
		"data":    snippet,
	})
}
