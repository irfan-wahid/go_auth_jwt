package databases

import (
	"fmt"
	"go_auth/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBs struct {
	Movies *gorm.DB
}

func PostgresConnection(dbName string, master config.DatabaseMasterConfig) *gorm.DB {
	dsnMaster := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s TimeZone=%s", master.Host, master.Port, master.DbUser, master.DbPassword, dbName, master.TimeZone)

	db, err := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("Error connection to database - %v", err)
	}

	log.Printf("Connection to database Postgresql %s success", dbName)

	return db
}

func MysqlConnection(dbName string, master config.DatabaseMasterConfig) *gorm.DB {
	dsnMaster := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", master.DbUser, master.DbPassword, master.Host, master.Port, dbName)

	db, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("Error connection to database - %v", err)
	}

	log.Printf("Connection to database Mysql %s success", dbName)

	return db
}

func (s *DBs) Init(master config.DatabaseMasterConfig) {
	if master.Dialect == "postgres" {
		s.Movies = PostgresConnection(master.DbName, master)
	} else if master.Dialect == "mysql" {
		s.Movies = MysqlConnection(master.DbName, master)
	}
}
