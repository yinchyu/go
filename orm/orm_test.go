package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
		db, _ := sql.Open("sqlite3", "gee.db")

		defer func() {

			_ = db.Close()

		}()
		_, _ = db.Exec("drop table if exists user;")

		_, _ = db.Exec("create table user(name text);")

		result, err := db.Exec("insert into user (`name`)values(?),(?)", "tom", "jack")
		if err != nil {
			affected, _ := result.RowsAffected()
			log.Println(affected)

		}

		row := db.QueryRow("select name from user limit 1;")
		var name string

		if err := row.Scan(&name); err == nil {
			log.Println(name)
		}

}

func TestEngine_NewSession(t *testing.T) {

	engine, _ := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)

}
