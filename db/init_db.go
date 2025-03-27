package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/internal"
	"github.com/zodiac182/tooth-health/server/model/system"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Pgsql struct {
	Username string
	Password string
	Dbname   string
	Port     string
	Path     string
}

var DB *gorm.DB

// TODO: template use
var (
	Path     = "localhost"
	Username = "postgres"
	Password = "password#123"
	Port     = "5432"
	Dbname   = "tooth_health"
)

// Dsn returns the data source name for the PostgreSQL connection.
func dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", Path, Username, Password, Dbname, Port)
}

func baseDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", Path, Username, Password, Port)
}

func InitDB() {

	logger.Debug("Initializing PostgreSQL database...")
	db, err := sql.Open("postgres", baseDsn())
	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL database: %v", err)
		return
	}
	defer db.Close()

	if !databaseExists(db, Dbname) {

		logger.Debug("Creating database...")
		err = createDatabase(db, Dbname)
		if err != nil {
			logger.Fatal("Failed to create database: %+v", err)
			return
		}

		logger.Debug("Database created successfully")
	}

	DB, err = gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL database: %+v", err)
		return
	}

	// 数据库自动迁移
	err = DB.AutoMigrate(
		&system.SysUser{},
		&system.CUser{},
		&system.TeethRecord{},
	)
	if err != nil {
		logger.Fatal("Failed to migrate database: %+v", err)
		return
	}

	// 创建初始用户
	createInitialUser()

	logger.Debug("PostgreSQL database initialized successfully")

}

func databaseExists(db *sql.DB, dbName string) bool {
	var exists bool

	err := db.QueryRow("SELECT 1 FROM pg_database WHERE datname = $1", dbName).Scan(&exists)
	if err != nil {
		logger.Info("Failed to check if database exists: %v", err)
		return false
	}
	return exists
}

func createDatabase(db *sql.DB, dbName string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	return err
}

func createInitialUser() {
	var count int64
	// 检查用户表是否已有记录
	err := DB.Model(&system.SysUser{}).Count(&count).Error
	if err != nil {
		logger.Fatal("Failed to count users: %v", err)
	}

	password, err := internal.HashPassword("password")
	if err != nil {
		logger.Fatal("Failed to hash password: %v", err)
		os.Exit(1)
	}

	if count == 0 {
		// 如果没有用户，创建初始用户
		initialUser := system.SysUser{
			Username: "admin",
			Password: password,
			Nickname: "管理员",
			Role:     system.AdminRole,
		}
		if err := DB.Create(&initialUser).Error; err != nil {
			logger.Fatal("Failed to create initial user: %v", err)
		}
		logger.Debug("Initial user 'admin' created successfully")
	}
}
