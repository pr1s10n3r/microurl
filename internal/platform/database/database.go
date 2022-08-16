package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var ErrConnectionNotOpen = errors.New("database connection has not been established")

type Connection struct {
	URL      string
	instance *sql.DB
}

func (c *Connection) Connect() error {
	db, err := sql.Open("mysql", c.URL)
	if err != nil {
		return err
	}
	c.instance = db

	return db.Ping()
}

func (c *Connection) Close() error {
	if c.instance == nil {
		return ErrConnectionNotOpen
	}

	return c.instance.Close()
}

func (c *Connection) DB() *sql.DB {
	return c.instance
}
