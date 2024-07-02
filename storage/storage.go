package storage

import (
	"encoding/json"
	"github.com/lil-zhi/go-comment/comment"
)

type Storage interface {
	Create(*comment.Comment) error
	Delete(*comment.Comment) error
	List(contentID, commentID, sourceType int) ([]*comment.Comment, error)
	Get(contentID, commentID int) (*comment.Comment, error)
}

type CommentGenerator struct {
	UserID      int
	Storage     Storage
	callBack    func() error
	deleteValid func() bool
	Recursive   bool
}

func NewCommentGenerator(userID int, storage Storage, recursive bool) *CommentGenerator {
	return &CommentGenerator{
		UserID:  userID,
		Storage: storage,
		callBack: func() error {
			return nil
		},
		deleteValid: func() bool {
			return true
		},
		Recursive: recursive,
	}
}

func (cg *CommentGenerator) SetCallBack(cb func() error) {
	cg.callBack = cb
}

func (cg *CommentGenerator) SetDeleteValid(dv func() bool) {
	cg.deleteValid = dv
}

func (cg *CommentGenerator) Generator(comment *comment.Comment) error {
	comment.UserID = cg.UserID
	if comment.CommentID != 0 {
		c, err := cg.Storage.Get(comment.ContentID, comment.CommentID)
		if err != nil {
			return err
		}
		if c.CommentFID != 0 {
			comment.CommentFID = c.CommentFID
		} else {
			comment.CommentFID = c.ID
		}
	}
	if err := cg.Storage.Create(comment); err != nil {
		return err
	}
	if err := cg.callBack(); err != nil {
		return err
	}
	return nil
}

func (cg *CommentGenerator) Delete(comment *comment.Comment) error {
	comment.UserID = cg.UserID
	if err := cg.Storage.Delete(comment); err != nil {
		return err
	}
	return nil
}

func (cg *CommentGenerator) GetComments(contentID, commentID int) ([]byte, error) {
	comments, err := cg.Storage.List(contentID, commentID, 1)
	if err != nil {
		return nil, err
	}
	if cg.Recursive {
		// 递归写法
	}
	var commentsRes []*comment.Comment
	var commentsMap = make(map[int][]*comment.Comment)
	for _, cur_comment := range comments {
		if cur_comment.CommentID == 0 {
			commentsRes = append(commentsRes, cur_comment)
		} else if len(commentsMap[cur_comment.CommentFID]) == 0 {
			commentsMap[cur_comment.CommentFID] = []*comment.Comment{cur_comment}
		} else {
			commentsMap[cur_comment.CommentFID] = append(commentsMap[cur_comment.CommentFID], cur_comment)
		}
	}
	for _, cur_comment := range commentsRes {
		cur_comment.Comments = commentsMap[cur_comment.ID]
		delete(commentsMap, cur_comment.ID)
	}
	for _, cur_comment := range commentsMap {
		commentsRes = append(commentsRes, &comment.Comment{Comments: cur_comment})
	}
	return json.Marshal(commentsRes)

}

func (cg *CommentGenerator) GetCommentsByComment(contentID, commentID int) (map[string]interface{}, error) {
	//return cg.Storage.get(id, 2)
	return nil, nil
}
