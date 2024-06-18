package services

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zerokkcoder/content-system/internal/process"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	goflow "github.com/s8sg/goflow/v1"
)

type CmsApp struct {
	db          *gorm.DB
	rdb         *redis.Client
	flowService *goflow.FlowService
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	// 连接数据库
	app.connDB()
	// 连接redis
	app.connRdb()
	app.flowService = flowService()
	go func() {
		process.ExecContentFlow(app.db)
	}()
	return app
}

func (app *CmsApp) connDB() {
	mysqlDB, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	mysqlDB = mysqlDB.Debug()

	app.db = mysqlDB
}

func flowService() *goflow.FlowService {
	fs := &goflow.FlowService{
		RedisURL: "localhost:6379",
	}
	return fs
}

func (app *CmsApp) connRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.rdb = rdb
}
