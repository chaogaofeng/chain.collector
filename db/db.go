package db

import (
	"fmt"
	"github.com/glodnet/chain.collector/schema"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MustNewDB(mode string, dsn string, opts ...gorm.Option) *gorm.DB {
	db, err := NewDB(mode, dsn, opts...)
	if err != nil {
		panic(err)
	}
	return db
}

func NewDB(mode string, dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	var conn gorm.Dialector
	switch mode {
	case "sqlite":
		conn = sqlite.Open(dsn)
	case "mysql":
		conn = mysql.Open(dsn)
	case "postgres":
		conn = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("invalid db engine. supported types: sqlite, mysql, postgres")
	}

	db, err := gorm.Open(conn, opts...)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		schema.Block{},
		schema.PreCommit{},
		schema.Transaction{},
		schema.PreCommit{},
		schema.Validator{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
