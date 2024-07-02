package comment

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	// 评论id
	ID int `json:"id" gorm:"primaryKey"`
	// 发布人id
	UserID int `json:"user_id" gorm:"Index:idx_user_id"`
	// 发布人name
	UserName string `json:"user_name" gorm:"-:migration"`
	// 内容id
	ContentID int `json:"content_id" gorm:"Index:idx_content_id"`
	// 某内容第一层评论id
	CommentFID int `json:"comment_fid" gorm:"column:comment_fid;Index:idx_comment_fid,priority:11"`
	// 某内容第一层评论id
	CommentID int `json:"comment_id" gorm:"column:comment_id;Index:idx_comment_id"`
	// 评论内容
	Content string `json:"content"`
	// 下面的评论
	Comments []*Comment `json:"comments" gorm:"-:migration"` // ignore this field when migrate with struct
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 删除时间
	DeleteAt gorm.DeletedAt `json:"delete_at"`
}

func (c *Comment) TableName() string {
	return "comments"
}
