package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/sdk_go/src/config"
	"github.com/sdk_go/src/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load the configuration file
	util.Load("config"+string(os.PathSeparator)+"config.json", config.Config)
	// createPrefences()

	fmt.Println(InsertOrder())
}

func createPrefences() {
	//params := "{\n    \"items\": [\n        {\n        \"title\": \"Dummy Item\",\n        \"description\": \"Multicolor Item\",\n        \"quantity\": 1,\n        \"currency_id\": \"ARS\",\n        \"unit_price\": 10.0\n        }\n    ]\n}"
	// controller.CreatePrefences()
}

//InsertOrder Insert the order
func InsertOrder() int64 {
	db, err := sql.Open("mysql", "root:root@/blueprint")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "use blueprint")

	stmt, _ := db.Prepare("insert into order_preferences(data_created, external_reference) values(?,?)")
	t := time.Now()
	res, error := stmt.Exec(t.Format("2006-01-02 15:04:05"), 1)
	// res, error := stmt.Exec("11", "1")

	if error != nil {
		panic(error)
	}
	id, _ := res.LastInsertId()

	return id
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}
