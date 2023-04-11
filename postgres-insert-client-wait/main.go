package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	connStr := "user=postgres dbname=test password=supersecret sslmode=disable application_name=testing"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	a := "pre-sleep"
	b := "no-read"
	c := fmt.Sprintf("%d", time.Now().Unix())

	//_, err = db.Exec(`INSERT INTO test(a, b, c) VALUES ($1, $2, $3)`, a, b, c)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("sleeping here")
	time.Sleep(60 * time.Second)

	a = "post-sleep"
	b = "read"
	c = fmt.Sprintf("%d", time.Now().Unix())

	_, err = db.Exec(`INSERT INTO test(a, b, c) VALUES ($1, $2, $3)`, a, b, c)
	if err != nil {
		panic(err)
	}

	fmt.Println("sleeping complete")

}
