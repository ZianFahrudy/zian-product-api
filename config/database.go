package config

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	"zian-product-api/common/exception"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	// port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")
	maxPoolOpen, err := strconv.Atoi(config.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPoolIdle, err := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))
	exception.PanicIfNeeded(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(username+":"+password+"@tcp("+host+")/"+dbName+"?parseTime=true"), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicIfNeeded(err)

	sqlDB, err := db.DB()
	exception.PanicIfNeeded(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	//autoMigrate
	//err = db.AutoMigrate(&entity.Product{})
	//err = db.AutoMigrate(&entity.Transaction{})
	//err = db.AutoMigrate(&entity.TransactionDetail{})
	//err = db.AutoMigrate(&entity.User{})
	//err = db.AutoMigrate(&entity.UserRole{})
	//exception.PanicLogging(err)
	return db
}
