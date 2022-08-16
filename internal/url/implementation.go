package url

import (
	"database/sql"

	"github.com/pr1s10n3r/microurl/internal/platform/database"
	"github.com/pr1s10n3r/microurl/pkg/url"
)

type urlRepoImpl struct {
	db *sql.DB
}

func NewUrlRepositoryImpl(conn database.Connection) urlRepoImpl {
	return urlRepoImpl{conn.DB()}
}

func (impl urlRepoImpl) Save(value string) (*url.URL, error) {
	stmt, err := impl.db.Prepare(`INSERT INTO urls (code, value) VALUES (?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	code := url.NewCode()
	if _, err = stmt.Exec(code, value); err != nil {
		return nil, err
	}

	return impl.FindByCode(code)
}

func (impl urlRepoImpl) FindByCode(code string) (*url.URL, error) {
	stmt, err := impl.db.Prepare(`SELECT id, value, created_at, updated_at FROM urls WHERE code = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	found := url.URL{Code: code}

	row := stmt.QueryRow(code)
	if err = row.Scan(&found.ID, &found.Value, &found.CreatedAt, &found.UpdatedAt); err != nil {
		return nil, err
	}

	return &found, nil
}
