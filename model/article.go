package model

import "gopkg.in/mgo.v2/bson"

/*
{
    id ObjectId,    *
    createtime timestemp,   *
    title string,   *
    tags []Tag{},   *
    comments [CommentId],
    content string, *
    description string,    *
    favorite [UserId],
    author User{},  *
    updatetime timestemp,   *
    meta {
        watch int,  *
        like int,   *
        num int *
    },  *
    keywords string,
    type int,    *    // 0: 私密文章, 1: 公开文章, 2: 简历, 3: 管理员介绍
	isdelete bool
}
*/
type Article struct {
	Id          bson.ObjectId   `bson:"_id" json:"id"`
	Createtime  int64           `bson:"createtime" json:"createtime"`
	Title       string          `bson:"title" json:"title"`
	Tags        []Tag           `bson:"tags" json:"tags"`
	Comments    []bson.ObjectId `bson:"comments" json:"comments"` //CommentId
	Content     string          `bson:"content" json:"content"`
	Description string          `bson:"description" json:"description"`
	Favorites   []bson.ObjectId `bson:"favorite" json:"favorite"` //UserId
	Author      User            `bson:"author" json:"author"`
	Updatetime  int64           `bson:"updatetime" json:"updatetime"`
	Meta        Meta            `bson:"meta" json:"meta"`
	Keywords    string          `bson:"keywords" json:"keywords"`
	Type        int             `bson:"type" json:"type"`
	Isdelete    bool            `bson:"isdelete" json:"isdelete"`
}
type Meta struct {
	Watch int `bson:"watch" json:"watch"`
	Like  int `bson:"like" json:"like"`
	Num   int `bson:"num" json:"num"`
}
