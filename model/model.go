package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitDB(connectMysqlStr string) error {
	var err error
	DB, err = gorm.Open("mysql", connectMysqlStr)
	if err != nil {
		logrus.Panic("Failed to connect mysql")
		return err
	}

	DB.SingularTable(true)
	DB.SetLogger(logrus.StandardLogger())

	// 修改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	return nil
}
