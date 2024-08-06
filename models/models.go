package models

import "time"

type BaseModel struct {
	ID        uint8 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// table pages
type Page struct {
	BaseModel
	Tity        string
	Body        string
	View        int
	IsPublished bool
}

// table posts
type Post struct {
	BaseModel
	Title        string     // title
	Body         string     // body
	View         int        // view count
	IsPublished  bool       // published or not
	Tags         []*Tag     `gorm:"-"` // tags of post
	Comments     []*Comment `gorm:"-"` // comments of post
	CommentTotal int        `gorm:"-"` // count of comment
}

// table tags
type Tag struct {
	BaseModel
	Name  string // tag name
	Total int    `gorm:"-"` // count of post
}

// table post_tags
type PostTag struct {
	BaseModel
	PostId uint // post id
	TagId  uint // tag id
}
