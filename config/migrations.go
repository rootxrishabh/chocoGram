package config

import (
	"fmt"

	"github.com/pressly/goose/v3"
)

func Migrate() {
    db, _ := GetDB().DB()

    if err := goose.SetDialect("mysql"); err != nil {
        fmt.Println("Error setting dialect")
        panic(err)
    }

    goose.Up(db, "db")
}