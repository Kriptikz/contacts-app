package main

import (
	"contacts-app/api"
	"log"
	"net/http"
	// "os"
	// "database/sql"
	// "fmt"
	// _ "github.com/libsql/libsql-client-go/libsql"
	// _ "modernc.org/sqlite"
)

func main() {
	// log.Println("Connecting to database")
	// var dbUrl = ProjectConfig.DatabaseUrl + "?authToken=" + ProjectConfig.DatabaseKey
	// db, err := sql.Open("libsql", dbUrl)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
	// 	os.Exit(1)
	// }

	// log.Println("Querying DB")

	// rows, err := db.Query("SELECT * from users")
	// if err != nil {
	// 	log.Fatal(fmt.Sprintf("Error selecting * from users: %v", err))
	// }
	// defer rows.Close()

	// columns, err := rows.Columns()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// values := make([]sql.RawBytes, len(columns))

	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// for rows.Next() {
	// 	err = rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	var value string
	// 	for i, col := range values {
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}
	// 		log.Println(columns[i], ": ", value)
	// 	}
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Starting server on port " + ProjectConfig.Port)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", api.RoutesHandler)
	http.ListenAndServe(":"+ProjectConfig.Port, nil)
}
