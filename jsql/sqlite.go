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
			Error: nil,
		},
		DbFile: "",
	}
}

func (s SqliteUtils) OpenFile(dbFile string) bool {

	s.DbFile = dbFile
	s.Db, s.Error = sql.Open("sqlite3", s.DbFile)

	return s.Error == nil

}

func (s SqliteUtils) OpenFileWithPass(dbFile string, user string, pass string) bool {

	s.DbFile = dbFile
	s.Db, s.Error = sql.Open("sqlite3", "file:"+s.DbFile+"?_auth&_auth_user="+user+"&_auth_pass="+pass)

	return s.Error == nil

}
