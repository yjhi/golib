package jsql

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Sql struct {
	DataBase *sql.DB
	IsOpen   bool
}

func NewSql() *Sql {

	s := &Sql{
		DataBase: nil,
		IsOpen:   false,
	}

	return s
}

func (s *Sql) CreateSqlServer(server string, user string, pass string, db string) bool {
	constr := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		server, user, pass, db)
	return s.ConnectSqlServer(constr)
}

func (s *Sql) ConnectSqlServer(constr string) bool {
	var err error
	s.DataBase, err = sql.Open("mssql", constr)

	if err != nil {
		s.IsOpen = true
	}

	return s.IsOpen
}

func (s *Sql) CloseSqlServer() {
	defer s.DataBase.Close()
}
