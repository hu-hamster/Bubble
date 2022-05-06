package dao

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hujing.ltd/buttle/model"
)

var DB *gorm.DB

func InitDB() {
	Collect()
	DB.AutoMigrate(&model.Todo{})
}

func Collect() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username, password, host, port, database, charset)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("连接数据库失败，错误信息：", err.Error())
	}
}

func Create(todo *model.Todo) error {
	d := DB.Create(todo)
	return d.Error
}

func GetAll() ([]model.Todo, error) {
	var todolist []model.Todo
	err := DB.Find(&todolist).Error
	if err != nil {
		return nil, err
	} else {
		return todolist, nil
	}
}

func Update(id string) error {
	var todo model.Todo
	err := DB.Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return err
	}
	todo.Status = !todo.Status
	err = DB.Save(&todo).Error
	return err
}

func Delete(id string) error {
	err := DB.Delete(&model.Todo{}, id).Error
	return err
}
