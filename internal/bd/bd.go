package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DataBase *sql.DB = nil

func InitDataBase() (err error) {
	if DataBase == nil {
		DataBase, err = sql.Open("sqlite3", "../source/db/copyposter.db")
		if err != nil {
			return fmt.Errorf("error open db %e", err)
		}
		if err := DataBase.Ping(); err != nil {
			return fmt.Errorf("error open db %e", err)
		}

		DataBase.Exec("create table test(t int)")
	}
	return err
}
