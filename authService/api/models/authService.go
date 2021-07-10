package models

import (
	"time"
)

// AuthInfo [...]
type AuthInfo struct {
	ID           int64     `gorm:"primaryKey;column:id;type:bigint;not null"`
	ClientName   string    `gorm:"column:client_name;type:varchar(256);not null;default:''"`   // 用户名称
	ClientSecret string    `gorm:"column:client_secret;type:varchar(256);not null;default:''"` // 客户端秘钥
	ClientID     string    `gorm:"column:client_id;type:varchar(256);not null;default:''"`     // 客户端名称
	CallbackURL  string    `gorm:"column:callback_url;type:varchar(512);not null;default:''"`  // 客户端回调地址
	IsDelete     bool      `gorm:"column:is_delete;type:tinyint(1);not null;default:1"`        // 是否删除
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

// AuthInfoColumns get sql column name.获取数据库列名
var AuthInfoColumns = struct {
	ID           string
	ClientName   string
	ClientSecret string
	ClientID     string
	CallbackURL  string
	IsDelete     string
	UpdateTime   string
	CreateTime   string
}{
	ID:           "id",
	ClientName:   "client_name",
	ClientSecret: "client_secret",
	ClientID:     "client_id",
	CallbackURL:  "callback_url",
	IsDelete:     "is_delete",
	UpdateTime:   "update_time",
	CreateTime:   "create_time",
}

// UserInfo [...]
type UserInfo struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint;not null"`
	Username   string    `gorm:"column:username;type:varchar(128);not null;default:''"`
	Password   string    `gorm:"column:password;type:varchar(128);not null;default:''"`
	Salt       string    `gorm:"column:salt;type:varchar(128);not null;default:''"`
	IsDelete   bool      `gorm:"column:is_delete;type:tinyint(1);not null;default:1"` // 是否删除
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

// UserInfoColumns get sql column name.获取数据库列名
var UserInfoColumns = struct {
	ID         string
	Username   string
	Password   string
	Salt       string
	IsDelete   string
	UpdateTime string
	CreateTime string
}{
	ID:         "id",
	Username:   "username",
	Password:   "password",
	Salt:       "salt",
	IsDelete:   "is_delete",
	UpdateTime: "update_time",
	CreateTime: "create_time",
}
