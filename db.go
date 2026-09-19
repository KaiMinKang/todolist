package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Warn),
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("MySQL数据库连接失败:%w", err)
	}
	if err = db.AutoMigrate(&UserModel{}, &TodoModel{}); err != nil {
		return nil, fmt.Errorf("建表失败: %w", err)
	}
	return db, nil
}
