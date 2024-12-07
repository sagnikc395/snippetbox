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
	//write the sql statement we want to execute.
	// split over 2 lines for readability
	stmt := `INSERT INTO snippets (title,content,created,expires)
	VALUES(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(),INTERVAL ? DAY))`

	//exec() method on embedded connection pool to execute the stmt

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	//use the LastInsertId() method on result object to get the ID our newly inserted
	// record in snippets table

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// return a specfici snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// retrun the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
