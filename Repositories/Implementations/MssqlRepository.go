package WWW

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type MssqlReposiory[T any] struct {
	connectionString string
	Db *sql.DB
	tableName string
}

func (rep *MssqlReposiory[T]) SetupRepository() {
	db, err := sql.Open("mssql", rep.connectionString)

	if err != nil {
		panic(err)
	}

	rep.Db = db
}

func (rep *MssqlReposiory[T]) CloseDb() {
	rep.Db.Close()
}

func (rep *MssqlReposiory[T]) GetAll() []T {
	rep.SetupRepository()
	defer rep.CloseDb()

	all, err := rep.Db.Query("SELECT * FROM %s", rep.tableName)
	if err != nil {
		panic(err)
	}

    var entities []T

	for all.Next() {
		var temp T
		all.Scan(&temp)
		append(entities, temp)
	}

	return entities
}