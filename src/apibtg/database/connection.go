package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type dbPgsql struct {
	dbPq *sql.DB
}

var (
	db *dbPgsql
)

// Initialize Connection
func InitConnection() error {
	var err error

	err = InitDbBTG()
	if err != nil {
		logger.Fatal(err, "Error in InitDbBTG")
		return err
	}
	return nil
}

// Closing Connection
func CloseConnection() {
	CloseDbBTG()
}

// Connect to DB BTG
func InitDbBTG() error {
	db = new(dbPgsql)
	server := viper.GetString("db-conn.server")
	port := viper.GetString("db-conn.port")
	user := viper.GetString("db-conn.user")
	pass := viper.GetString("db-conn.pass")
	schema := viper.GetString("db-conn.schema")

	return initConnectionPq(db, server, port, user, pass, schema)

}
func CloseDbBTG() {
	logger.Println("Closing BTG_Test connection...")
	db.CloseConnectionPq()
}

func GetConnectionBTG() (*sql.DB, error) {
	return db.GetConnectionPq()
}

// General Connection
func initConnectionPq(dbCon *dbPgsql, server, port, user, pass, scheme string) error {
	var db *sql.DB
	db, err := createConnectionPq(server, port, user, pass, scheme)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	dbCon.dbPq = db
	return nil
}

func (dpq *dbPgsql) CloseConnectionPq() {
	dpq.dbPq.Close()
}

func (dpq *dbPgsql) GetConnectionPq() (*sql.DB, error) {
	if dpq.dbPq == nil {
		return nil, errors.New("pgsql: failed to get connection")
	}

	if err := dpq.dbPq.Ping(); err != nil {
		logger.Fatal(err, "Error in PingDB PGSQL")
		return nil, err
	}
	return dpq.dbPq, nil
}

func createConnectionPq(server, port, user, pass, schema string) (*sql.DB, error) {
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server,
		port,
		user,
		pass,
		schema)
	conn, errdb := sql.Open("postgres", connInfo)
	if errdb != nil {
		logger.Fatal(errdb, "Error in OpenConnection PgSQL")
		return nil, errdb
	}
	if err := conn.Ping(); err != nil {
		logger.Fatal(err, "Error in PingConnection PgSQL")
		return nil, err
	}
	return conn, nil
}
