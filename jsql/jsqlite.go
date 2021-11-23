package jsql
import (
	"database/sql"
	"fmt"

	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteUtils struct {
	DbFile string
	Error  error
	Db     *sql.DB
	Result sql.Result
	Tx     *sql.Tx
	Stmt   *sql.Stmt
}

func (s *SqliteUtils) OpenFile(dbFile string) bool {

	s.DbFile = dbFile
	s.Db, s.Error = sql.Open("sqlite3", s.DbFile)

	return s.Error == nil

}

func (s *SqliteUtils) OpenFileWithPass(dbFile string, user string, pass string) bool {

	s.DbFile = dbFile
	s.Db, s.Error = sql.Open("sqlite3", "file:"+s.DbFile+"?_auth&_auth_user="+user+"&_auth_pass="+pass)

	return s.Error == nil

}

func (s *SqliteUtils) ExecCmd(cmd string) bool {

	if s.Db == nil {
		return false
	}

	var stmt *sql.Stmt
	stmt, s.Error = s.Db.Prepare(cmd)

	if s.Error != nil {
		return false
	}
	s.Result, s.Error = stmt.Exec()

	return s.Error == nil
}

func (s *SqliteUtils) ExecCmdWithArgs(cmd string, args ...interface{}) bool {

	if s.Db == nil {
		return false
	}

	var stmt *sql.Stmt
	stmt, s.Error = s.Db.Prepare(cmd)

	if s.Error != nil {
		return false
	}

	s.Result, s.Error = stmt.Exec(args...)

	return s.Error == nil
}

func (s *SqliteUtils) GetInt(cmd string) (int64, bool) {

	if s.Db == nil {
		return 0, false
	}

	var count int64

	s.Error = s.Db.QueryRow(cmd).Scan(&count)

	if s.Error != nil {
		return 0, false
	}
	return count, true
}

func (s *SqliteUtils) GetString(cmd string) (string, bool) {

	if s.Db == nil {
		return "", false
	}

	var str string

	s.Error = s.Db.QueryRow(cmd).Scan(&str)

	if s.Error != nil {
		return "", false
	}

	return str, true
}

func (s *SqliteUtils) Begin() bool {

	if s.Db == nil {
		return false
	}

	s.Tx, s.Error = s.Db.Begin()

	return s.Error == nil

}

func (s *SqliteUtils) End() bool {
	if s.Db == nil {
		return false
	}

	if s.Tx == nil {
		return false
	}

	s.Error = s.Tx.Commit()

	if s.Error != nil {
		s.Tx.Rollback()

		return false
	}

	return true

}

func (s *SqliteUtils) Prepare(cmd string) bool {

	if s.Db == nil {
		return false
	}

	s.Stmt, s.Error = s.Db.Prepare(cmd)

	return s.Error == nil
}

func (s *SqliteUtils) ExecBatch(args ...interface{}) bool {
	if s.Db == nil {
		return false
	}

	if s.Stmt == nil {
		return false
	}

	s.Result, s.Error = s.Stmt.Exec(args...)

	return s.Error == nil

}