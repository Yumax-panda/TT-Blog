package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserInfo interface {
	GetID() uuid.UUID
	GetName() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	IsProfileAvailable() bool
}

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Name      string    `gorm:"type:varchar(32);not null"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`

	Profile *UserProfile `gorm:"constraint:user_profiles_user_id_users_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (user *User) TableName() string {
	return "users"
}

type UserProfile struct {
	UserID    uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Bio       string    `gorm:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}

type ExternalProviderUser struct {
	UserID       uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	ProviderName string    `gorm:"type:varchar(30);not null;primaryKey;uniqueIndex:idx_external_provider_users_provider_name_external_id,priority:1"`
	ExternalID   string    `gorm:"type:varchar(100);not null;uniqueIndex:idx_external_provider_users_provider_name_external_id,priority:2"`
	Extra        JSON      `gorm:"type:text;not null"`
	CreatedAt    time.Time `gorm:"precision:6"`
	UpdatedAt    time.Time `gorm:"precision:6"`

	User *User `gorm:"constraint:external_provider_users_user_id_users_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ExternalProviderUser) TableName() string {
	return "external_provider_users"
}

func (user *User) GetID() uuid.UUID {
	return user.ID
}

func (user *User) GetName() string {
	return user.Name
}

func (user *User) GetCreatedAt() time.Time {
	return user.CreatedAt
}

func (user *User) GetUpdatedAt() time.Time {
	if user.IsProfileAvailable() {
		if user.Profile.UpdatedAt.After(user.UpdatedAt) {
			return user.Profile.UpdatedAt
		}
	}
	return user.UpdatedAt
}

func (user *User) IsProfileAvailable() bool {
	return user.Profile != nil
}
