package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Org struct {
	org_code string
	org_name string
}

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBHOST") + ":3306"
	cfg.DBName = "incent_datamart"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	orgs, err := Orgs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", orgs)

}

func Orgs() ([]Org, error) {
	var orgs []Org
	rows, err := db.Query("SELECT * FROM org_hierarchy")
	if err != nil {
		return nil, fmt.Errorf("Orgs: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var scanned_org Org
		if err := rows.Scan(&scanned_org.org_code, &scanned_org.org_name); err != nil {
			return nil, fmt.Errorf("Orgs: %v", err)
		}
		orgs = append(orgs, scanned_org)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Orgs: %v", err)
	}
	return orgs, nil
}
