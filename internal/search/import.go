package search

import (
	_ "github.com/NewAlist/alist/internal/search/bleve"
	_ "github.com/NewAlist/alist/internal/search/db"
	_ "github.com/NewAlist/alist/internal/search/db_non_full_text"
	_ "github.com/NewAlist/alist/internal/search/meilisearch"
)
