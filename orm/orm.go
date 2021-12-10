package main

import (
	"database/sql"
	"orm/dialect"
	"orm/log"
	"orm/session"
)

type Engine struct {
	// 就封装 database 中的DB字段
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	// ping 好手段， 如果没有问题的话就不返回error
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// 创建的时候也需要  dialect
	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dialect}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
