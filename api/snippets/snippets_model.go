package snippets

//Snippet describes a code snippet
type Snippet struct {
	SnippetID  uint   `json:"snippetID"`
	Snippet    string `json:"snippet" binding:"required"`
	LanguageID uint   `json:"languageID" binding:"required"`
}
