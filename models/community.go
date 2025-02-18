package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"commid,string" db:"community_id" binding:"required"`
	Name         string    `json:"commname" db:"community_name" binding:"required"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}

type CommPosts struct {
	Community *Community
	ID        int64 `json:"id,string" db:"post_id"`
	AuthorID  int64 `json:"author_id,string" db:"author_id"`
	//CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
	Status int32  `json:"status,string" db:"status"`
	Title  string `json:"title" db:"title" binding:"required"`
	//Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
