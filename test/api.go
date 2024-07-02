package main

import (
	"errors"
	"github.com/lil-zhi/go-comment/comment"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
	//storage2.Storage
}

var storage Storage

func init() {
	var err error
	storage.db, err = NewDB()
	if err != nil {
		panic(err)
	}
}

func (s *Storage) Create(c *comment.Comment) error {
	return s.db.Omit("user_name").Create(c).Error
}

func (s *Storage) Delete(c *comment.Comment) error {
	return s.db.Where("id = ?", c.ID).Delete(&c).Error
}

func (s *Storage) List(contentID, commentID, sourceType int) ([]*comment.Comment, error) {
	var comments []*comment.Comment
	if sourceType == 1 {
		err := s.db.Where("content_id = ?", contentID).Find(&comments).Error
		return comments, err
	}
	if sourceType == 2 {
		err := s.db.Where("content_id = ? and comment_fid = ?", contentID, commentID).Find(&comments).Error
		return comments, err
	}
	return nil, errors.New("source type is not supported")
}

func (s *Storage) Get(contentID, commentID int) (*comment.Comment, error) {
	var c *comment.Comment
	err := s.db.Where("content_id = ? and id = ?", contentID, commentID).First(&c).Error
	return c, err
}
