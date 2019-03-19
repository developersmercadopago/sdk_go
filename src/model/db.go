package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sdk_go/src/config"
	//mysql import
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

//InsertOrder Insert the order
func InsertOrder(external_reference string) int64 {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", config.Config.DataBaseUser, config.Config.DataBasePass, config.Config.DataBaseName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "use blueprint")

	stmt, _ := db.Prepare("insert into order_preferences(data_created, external_reference) values(?,?)")
	t := time.Now()
	res, error := stmt.Exec(t.Format("2006-01-02 15:04:05"), 1)

	if error != nil {
		panic(error)
	}

	id, _ := res.LastInsertId()
	fmt.Println(id)
	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)

	return id
}
