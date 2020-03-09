package database

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-echo-api/models"
	"time"
)

func RegisterTxDB(name string) {
	txdb.Register(name, "postgres", "postgres://postgres:postgres@localhost:5433/echo_api?sslmode=disable")
}

// PrepareTestDB prepare test DB according to txdb name
func PrepareTestDB(withName string) (*gorm.DB, error) {

	sqlDB, err := sql.Open(withName, fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	db, err := gorm.Open("postgres", sqlDB)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models.User{})
	return db, err
}

// CleanTestDB drops all tables from test DB
func CleanTestDB(db *gorm.DB) {
	_ = db.Close()
}
