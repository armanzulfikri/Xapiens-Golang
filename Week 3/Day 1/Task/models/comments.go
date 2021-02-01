package models

//comments Model

type Comments struct {
	PostId uint   `json:"postId`
	ID     uint   `json:id`
	Name   string `json:name`
	Email  string `json:email`
	Body   string `json:body`
}
