package db

import (
	"context"
	"database/sql"
	"time"

	"role-based-access-control-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Conn  *gorm.DB
	sqlDB *sql.DB
}

func Connect(ctx context.Context, databaseURL string) (*DB, error) {
	gormDB, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(pingCtx); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	return &DB{Conn: gormDB, sqlDB: sqlDB}, nil
}

func (d *DB) Close() error {
	if d == nil || d.sqlDB == nil {
		return nil
	}
	return d.sqlDB.Close()
}

func (d *DB) AutoMigrate() error {
	if d == nil || d.Conn == nil {
		return nil
	}
	return d.Conn.AutoMigrate(&models.User{})
}
