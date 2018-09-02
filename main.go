package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BrunoBA/go_lang_test/api/models/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "root:@/go_lang_test")
	db, err := sql.Open("mysql", "bfxj8igiok2x69ut:kktf2byp3rlpou95@nivk0hz7m5elq4ql.chr7pe7iynqr.eu-west-1.rds.amazonaws.com:3306/agcltfrrpw681ogm")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, first_name FROM users")
	defer rows.Close()

	for rows.Next() {
		var u = user.NewUser()

		rows.Scan(&u.Id, &u.Name)

		fmt.Println(u)
	}
}
