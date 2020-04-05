package snippets

import (
	"html"

	"github.com/zcericola/hexmouse-backend/api/utils"

	"github.com/zcericola/hexmouse-backend/db"
)

//EscapeSpecialCharacters prepares a string snippet to be stored in the database
func EscapeSpecialCharacters(snippet string) string {
	formattedStr := html.EscapeString(snippet)
	return formattedStr
}

//UnescapeSpecialCharacters reverses the string escapement
func UnescapeSpecialCharacters(snippet string) string {
	formattedStr := html.UnescapeString(snippet)
	return formattedStr
}

//LinkSnippetToUser ties the snippet to a specific userID
func LinkSnippetToUser(userID uint, snippetID uint) {
	insertStatement := `
		INSERT INTO users_snippets
		(user_id, snippet_id)
		VALUES($1, $2)
	`
	stmt, err := db.DB.Prepare(insertStatement)
	utils.HandleError(err)
	_, err = stmt.Exec(
		userID,
		snippetID,
	)
	return
}
