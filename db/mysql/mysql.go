package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	User string
	Pass string
	Host string
	Name string
}

var pool = make(map[string]*gorm.DB)

func GetDB(name string) *gorm.DB {
	db := pool[name]
	if db == nil {
		panic(fmt.Sprintf("db %s not exit", name))
	}
	return db
}

func InitMysql(dbs []DB) error {
	for _, db := range dbs {
		dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.Pass, db.Host, db.Name)
		gdb, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}))
		if err != nil {
			return err
		}
		pool[db.Name] = gdb
	}
	return nil
}
