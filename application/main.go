package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Username string    `gorm:"size:50"`
	Password string    `gorm:"type:varchar(100)"`
	Token    string
	gorm.Model
}

func GenerateUuid() uuid.UUID {
	UserId, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return uuid.Nil
	}
	return UserId
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

}
