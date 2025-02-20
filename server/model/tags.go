package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tag struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Name      string    `gorm:"type:VARCHAR(30) COLLATE utf8mb4_bin NOT NULL;uniqueIndex:name"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (*Tag) TableName() string {
	return "tags"
}

type PostTag interface {
	GetPostID() uuid.UUID
	GetTagID() uuid.UUID
	GetTag() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type PostsTag struct {
	PostID    uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	TagID     uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	CreatedAt time.Time `gorm:"precision:6;index"`
	UpdatedAt time.Time `gorm:"precision:6"`

	Post *Post `gorm:"constraint:posts_tags_post_id_posts_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tag  Tag   `gorm:"constraint:posts_tags_tag_id_tags_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (*PostsTag) TableName() string {
	return "posts_tags"
}

func (p *PostsTag) GetPostID() uuid.UUID {
	return p.PostID
}

func (p *PostsTag) GetTagID() uuid.UUID {
	return p.TagID
}

func (p *PostsTag) GetTag() string {
	return p.Tag.Name
}

func (p *PostsTag) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *PostsTag) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}
