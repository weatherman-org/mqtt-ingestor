package util

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostgresPool(dsn string) *pgxpool.Pool {
	count := 0
	for {
		db, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			count++
		} else {
			fmt.Println("connected to postgres!")
			return db
		}
		if count == 5 {
			fmt.Println("unable to connect to postgres...", err)
			fmt.Println("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}

func CreateDatabase(conn *pgxpool.Pool) {
	count := 0
	for {
		query := `CREATE DATABASE "telemetry";`
		_, err := conn.Exec(context.Background(), query)
		if err != nil {
			if e, ok := err.(*pgconn.PgError); ok && e.Code == "42P04" {
				fmt.Println("database already exists...")
				return
			}
			count++
		} else {
			fmt.Println("database created!")
			return
		}
		if count == 5 {
			fmt.Println("unable to create database...", err)
			fmt.Println("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}
