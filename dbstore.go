package ihttp

import (
	"database/sql"
	"fmt"
	"github.com/gitkeng/ihttp/log"
	"github.com/gitkeng/ihttp/util/stringutil"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

type IDBStore interface {
	// Conn return dbConn to database
	Conn() *sql.DB
	// Close close database dbConn
	Close() error
	// Config return database config
	Config() IDBConfig
}

type DBStore struct {
	dbConn *sql.DB
	config IDBConfig
}

func NewDBStore(cfg IDBConfig) (IDBStore, error) {
	retries := -1
	for {
		retries++
		if retries > cfg.GetRetryLimits()-1 {
			return nil, fmt.Errorf("create dbConn retry exceed limits")
		}
		var conn *sql.DB
		var err error
		switch cfg.GetProvider() {
		case POSTGRES:
			conn, err = pqOpen(cfg.GetURL(), cfg.GetUser(), cfg.GetPassword(), cfg.GetDatabaseName())
		case MYSQL:
			conn, err = mySqlOpen(cfg.GetURL(), cfg.GetUser(), cfg.GetPassword(), cfg.GetDatabaseName())
		}

		if err != nil {
			log.Warnf("database context name %s connect fail with err %s", cfg.GetContextName(), err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		if conn != nil {
			conn.SetConnMaxLifetime(time.Second * time.Duration(cfg.GetConnectionMaxLifeTime()))
			conn.SetMaxIdleConns(cfg.GetMaxIdleConns())
			conn.SetMaxOpenConns(cfg.GetMaxOpenConns())
			//ping for check err
			err = conn.Ping()
			if err != nil {
				log.Warnf("database context name %s ping fail with err %s", cfg.GetContextName(), err.Error())
				time.Sleep(time.Second * 1)
				continue
			}
		} else {
			log.Warnf("database context name %s ping fail not found dbConn", cfg.GetContextName())
			time.Sleep(time.Second * 1)
			continue
		}

		return &DBStore{
			dbConn: conn,
			config: cfg,
		}, nil

	}
}

func (db *DBStore) Conn() *sql.DB {
	return db.dbConn

}

func (db *DBStore) Close() error {
	return db.dbConn.Close()
}

func (db *DBStore) Config() IDBConfig {
	return db.config
}

func pqOpen(url string, user string, password string, dbName string) (*sql.DB, error) {
	if stringutil.IsEmptyString(url) {
		return nil, ErrDBURLRequire
	}
	if stringutil.IsEmptyString(user) {
		return nil, ErrDBUserRequire
	}
	if stringutil.IsEmptyString(password) {
		return nil, ErrDBPasswordRequire
	}
	if stringutil.IsEmptyString(dbName) {
		return nil, ErrDBNameRequire
	}

	urls := strings.Split(url, ":")
	if len(urls) != 2 {
		return nil, ErrDBURLPattern(url)
	}

	urlStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, url, dbName)
	db, err := sql.Open("postgres", urlStr)
	if err != nil {
		return nil, err
	}
	return db, nil

}

func mySqlOpen(url string, user string, password string, dbName string) (*sql.DB, error) {
	if stringutil.IsEmptyString(url) {
		return nil, ErrDBURLRequire
	}
	if stringutil.IsEmptyString(user) {
		return nil, ErrDBUserRequire
	}
	if stringutil.IsEmptyString(password) {
		return nil, ErrDBPasswordRequire
	}
	if stringutil.IsEmptyString(dbName) {
		return nil, ErrDBNameRequire
	}

	urlStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&multiStatements=true", user, password, url, dbName)
	db, err := sql.Open("mysql", urlStr)
	if err != nil {
		return nil, err
	}
	return db, nil

}
