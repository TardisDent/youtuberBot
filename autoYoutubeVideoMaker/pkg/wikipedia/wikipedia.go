package wikipedia

func GetArticle(title string) (*Article, error) {
	// Get the article from Wikipedia.
	article, err := wikipedia.NewArticle(title)
	if err != nil {
		return nil, err
	}

	// Return the article.
	return article, nil
}

type Article struct {
	Title         string
	Content       string
	Images        []string
	Categories    []string
	References    []string
	ExternalLinks []string
}
 