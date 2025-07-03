package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlConnection() (*sql.DB, error) {

	b2c := connMySqlB2C()
	conn, err := connectMySQL(b2c)
	if err != nil {
		return nil, err
	}

	setupConnectionPool(conn)
	return conn, nil
}

func NewMySqlBcloudMS() (*sql.DB, error) {

	ms := connMySqlBcloudMS()
	conn, err := connectMySQL(ms)
	if err != nil {
		return nil, err
	}

	setupConnectionPool(conn)
	return conn, nil
}

func connectMySQL(config mysql.Config) (*sql.DB, error) {

	conn, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func connMySqlB2C() mysql.Config {

	config := mysql.Config{
		User:                 "censored",
		Passwd:               "censored",
		Net:                  "tcp",
		Addr:                 "amazonaws.com",
		AllowNativePasswords: true,
		ParseTime:            true,
		DBName:               "spock",
	}

	return config
}

func connMySqlBcloudMS() mysql.Config {

	config := mysql.Config{
		User:                 "censored",
		Passwd:               "censored",
		Net:                  "tcp",
		Addr:                 "amazonaws.com",
		AllowNativePasswords: true,
		ParseTime:            true,
		DBName:               "spock",
	}

	return config
}

func setupConnectionPool(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)
}
