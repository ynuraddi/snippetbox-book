package sqlite

import (
	"database/sql"

	"github.com/Tlepkali/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	var stmt string
	if expires == "1" {
		stmt = `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, CURRENT_TIMESTAMP, DATE(CURRENT_TIMESTAMP, '+1 DAY'))`
	} else if expires == "7" {
		stmt = `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, CURRENT_TIMESTAMP, DATE(CURRENT_TIMESTAMP, '+7 DAYS'))`
	} else {
		stmt = `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, CURRENT_TIMESTAMP, DATE(CURRENT_TIMESTAMP, '+1 YEAR'))`
	}

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}

	err := m.DB.QueryRow("SELECT * FROM snippets WHERE id = ?", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
