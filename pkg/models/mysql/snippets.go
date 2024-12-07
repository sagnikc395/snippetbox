package mysql

import (
	"database/sql"

	"github.com/sagnikc395/snippetbox/pkg/models"
)

// define a snippetmodel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into the databse
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// return a specfici snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// retrun the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
