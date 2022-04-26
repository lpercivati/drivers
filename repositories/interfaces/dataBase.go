package repositories

import "database/sql"

type DataBaseRepository interface {
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
}
