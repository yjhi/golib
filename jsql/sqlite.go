package jsql

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
*  add by yjh 211125
*  for sqlserver
*
 */
type SqliteUtils struct {
	SqlUtils
	DbFile string
}

func BuildSqlite() *SqliteUtils {
	return &SqliteUtils{
		SqlUtils: SqlUtils{
			_error: nil,
		},
		DbFile: "",
	}
}

func CreateSqliteString(file string, config string) string {

	s := "file:" + file

	if len(config) > 0 {
		s += "?" + config
	}

	return s
}

func CreateSqliteStringWithPass(file string, user string, pass string, config string) string {
	s := "file:" + file + "?_auth&_auth_user=" + user + "&_auth_pass=" + pass

	if len(config) > 0 {
		s += "&" + config
	}
	return s

}

func (s *SqliteUtils) OpenFile(dbFile string) bool {

	s.DbFile = dbFile
	s.Db, s._error = sql.Open("sqlite3", s.DbFile)

	return s._error == nil

}

func (s *SqliteUtils) OpenFileWithPass(dbFile string, user string, pass string) bool {

	s.DbFile = dbFile
	s.Db, s._error = sql.Open("sqlite3", CreateSqliteStringWithPass(dbFile, user, pass, ""))

	return s._error == nil

}
