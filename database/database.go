package database

import (
	"almcm.poscoict.com/scm/pme/curly-engine/configure"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DbInfo struct {
	Db       *gorm.DB
	Id       string
	Password string
	Ip       string
	Port     int
	DbName   string
}

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "sky"
//	password = "1234"
//	dbname   = "postgres"
//)

func (dbInfo *DbInfo) GetConnection() (db *gorm.DB, err error) {

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)

	//db, err := gorm.Open("postgres", psqlInfo)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbInfo.Ip, dbInfo.Id, dbInfo.Password, dbInfo.DbName, dbInfo.Port)
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dbInfo.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")
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

func ConnNewDbFromConfig() (dbInfo *DbInfo, err error) {
	conf := configure.GetConfig()

	dbInfo = &DbInfo{
		Id:       conf.Db.UserId,
		Password: conf.Db.Password,
		Ip:       conf.Db.IpAddress,
		Port:     conf.Db.Port,
		DbName:   conf.Db.DbName,
	}

	_, err = dbInfo.ConnectDatabase()
	if err != nil {
		_, err = dbInfo.GetConnection()
		if err != nil {
			Loge("Failed to connect Database %s : %s", dbInfo.DbName, err)
			return nil, err
		}
		dbInfo.CreateDatabase()
	}
	return dbInfo, nil
}
