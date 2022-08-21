package mysql

import (
	"database/sql"
	_ "database/sql/driver"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Load wrapped connection
func Load(user, pass, host, port, database string) (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?collation=utf8mb4_unicode_ci&parseTime=true")
	if err != nil {
		return nil, err
	}

	// Limit resource usage
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(30 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
