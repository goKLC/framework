package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goKLC/goKLC"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Surname   string    `json:"surname" gorm:"type:varchar(100)"`
	Username  string    `json:"username" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique_index;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func init() {
	app := goKLC.GetApp()
	app.DB().AutoMigrate(&User{})
	app.Auth().Model = &User{}
}
