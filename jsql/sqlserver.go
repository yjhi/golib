package jsql

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
*  add by yjh 211125
*  for sqlserver
*
 */
type SqlServerUtils struct {
	SqlUtils
	IsOpen bool
}

func BuildSqlServer() *SqlServerUtils {

	s := &SqlServerUtils{
		SqlUtils: SqlUtils{
			Error: nil,
		},
		IsOpen: false,
	}

	return s
}

func CreateSqlServerString(server string, user string, pass string, db string) string {
	constr := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		server, user, pass, db)
	return constr
}

func (s *SqlServerUtils) ConnectSqlServerString(server string, user string, pass string, db string) bool {
	constr := CreateSqlServerString(server, user, pass, db)

	return s.ConnectSqlServer(constr)
}

func (s *SqlServerUtils) ConnectSqlServer(constr string) bool {

	s.Db, s.Error = sql.Open("mssql", constr)

	if s.Error != nil {
		s.IsOpen = true
	}

	return s.IsOpen
}

func (s *SqlServerUtils) CloseSqlServer() {
	defer s.Db.Close()
}
