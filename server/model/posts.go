package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:char(36);not null;primaryKey"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;"`
	Content   string         `gorm:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	CreatedAt time.Time      `gorm:"precision:6;"`
	UpdatedAt time.Time      `gorm:"precision:6;"`
	DeletedAt gorm.DeletedAt `gorm:"precision:6;"`

	User *User `gorm:"constraint:posts_user_id_users_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (p Post) TableName() string {
	return "posts"
}
