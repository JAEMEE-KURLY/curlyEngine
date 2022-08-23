package gormdb

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"time"

	"almcm.poscoict.com/scm/pme/curly-engine/configure"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInfo struct {
	Db        *gorm.DB
	Id        string
	Password  string
	Ip        string
	Port      int
	DbName    string
	DbNameLog string
	DbNameHmi string
}

var (
	MainDB *DbInfo
	LogDB  *DbInfo
)

func InitSingletonDB() error {
	conf := configure.GetConfig()

	if MainDB == nil {
		MainDB = &DbInfo{
			Id:       conf.Db.UserId,
			Password: conf.Db.Password,
			Ip:       conf.Db.IpAddress,
			Port:     conf.Db.Port,
			DbName:   conf.Db.DbName,
		}
		err := ConnectDatabase(MainDB)
		if err != nil {
			return err
		}
	}

	if LogDB == nil {
		LogDB = &DbInfo{
			Id:        conf.Db.UserId,
			Password:  conf.Db.Password,
			Ip:        conf.Db.IpAddress,
			Port:      conf.Db.Port,
			DbName:    conf.Db.DbName,
			DbNameLog: conf.Db.DbNameLog,
		}
		err := ConnectDatabase(LogDB)
		if err != nil {
			return err
		}
	}

	Logd("Success Init Singleton DB")

	return nil
}
func (dbInfo *DbInfo) GetConnection() (db *gorm.DB, err error) {
	var dsn string

	if dbInfo.Port < 0 {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbInfo.Id, dbInfo.Password, dbInfo.Ip, dbInfo.DbName)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul", dbInfo.Ip, dbInfo.Id, dbInfo.Password, dbInfo.DbName)
		//dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		//	dbInfo.Id, dbInfo.Password, dbInfo.Ip, dbInfo.Port, dbInfo.DbName)
	}

	dbInfo.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(GetLogWriter(), "\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Warn,
				Colorful:      true,
			}),
	})
	if err != nil {
		Loge("Failed to connect database %s : %s", dbInfo.DbName, err)
		return nil, err
	}
	return dbInfo.Db, nil
}

func (dbInfo *DbInfo) ConnectDatabase() (db *gorm.DB, err error) {
	var dsn string

	if dbInfo.Port < 0 {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbInfo.Id, dbInfo.Password, dbInfo.Ip, dbInfo.DbName)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul", dbInfo.Ip, dbInfo.Id, dbInfo.Password, dbInfo.DbName)
		//dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		//	dbInfo.Id, dbInfo.Password, dbInfo.Ip, dbInfo.Port, dbInfo.DbName)
	}

	dbInfo.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(GetLogWriter(), "\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Warn,
				Colorful:      true,
			}),
	})
	if err != nil {
		Loge("Failed to connect database %s : %s", dbInfo.DbName, err)
		return nil, err
	}
	return dbInfo.Db, nil
}
func (dbInfo *DbInfo) CreateDatabase() {
	if dbInfo.Db == nil || len(dbInfo.DbName) < 1 {
		return
	}
	dbInfo.Db.Exec("CREATE DATABASE IF NOT EXISTS " + dbInfo.DbName)
	dbInfo.Db.Exec("commit;")
}
func ConnectDatabase(dbInfo *DbInfo) (err error) {
	_, err = dbInfo.ConnectDatabase()
	if err != nil {
		_, err = dbInfo.GetConnection()
		if err != nil {
			Loge("Failed to connect Database : %s", err)
			return err
		}
		dbInfo.CreateDatabase()
		_, err = dbInfo.ConnectDatabase()
		if err != nil {
			Loge("Error: Create DB but failed to connect DB")
			return err
		}
	}
	return nil
}
