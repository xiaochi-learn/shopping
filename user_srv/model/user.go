package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32          `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

/*
1、密码密文保存
2、不可反解
*/

type User struct {
	BaseModel
	Mobel    string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"type:varchar(8);column:gender;default:'male';comment:'female表示女,male表示男'"`
	Role     int        `gorm:"type:int;default:1;type:int;comment:'1表示普通用户,2表示管理员'"`
}
