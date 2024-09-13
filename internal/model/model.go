package model

import (
	"LiuYanXiBlog/global"
	"LiuYanXiBlog/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"

	//"github.com/go-sql-driver/mysql"
	//"gorm.io/driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	DeletedOn  string `json:"deleted_on"`
	IsDel      string `json:"is_del"`
}

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// 构建数据库连接字符串
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn := fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 设置日志级别
	if global.ServerSetting.RunMode == "debug" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略记录未找到错误
				Colorful:                  true,        // 彩色打印
			},
		)
		db = db.Set("gorm:logger", newLogger)
	}

	// 设置全局表名策略为单数形式
	db = db.Set("gorm:naming_strategy", struct {
		SingularTable bool
	}{SingularTable: true})

	// 设置数据库连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleconns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(databaseSetting.ConnMaxLifetimeHours) * time.Hour)

	return db, nil
}
