package main

import (
	"fmt"
	"github.com/lil-zhi/go-comment/comment"
	storage2 "github.com/lil-zhi/go-comment/storage"
)

func main() {
	err := storage.db.AutoMigrate(comment.Comment{})
	if err != nil {
		panic(err)
	}
	//fmt.Println(storage)
	cg := storage2.NewCommentGenerator(111, &storage, true)
	err = cg.Generator(&comment.Comment{
		ContentID: 1,
		CommentID: 11,
		Content:   "我是第二层回复",
	})
	if err != nil {
		panic(err)
	}
	comments, err := cg.GetComments(1, 1)
	fmt.Println(string(comments), err)
}
