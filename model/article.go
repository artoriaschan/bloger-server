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
	Tags        []bson.ObjectId `bson:"tags" json:"tags"`             // tagId
	Categories  []bson.ObjectId `bson:"categories" json:"categories"` // cateId
	Comments    []bson.ObjectId `bson:"comments" json:"comments"`     // CommentId
	Content     string          `bson:"content" json:"content"`
	Description string          `bson:"description" json:"description"`
	Favorites   []bson.ObjectId `bson:"favorite" json:"favorite"` // UserId
	Author      bson.ObjectId   `bson:"author" json:"author"`
	Updatetime  int64           `bson:"updatetime" json:"updatetime"`
	Meta        Meta            `bson:"meta" json:"meta"`
	Cover       string          `bson:"cover" json:"cover"`
	Keywords    string          `bson:"keywords" json:"keywords"`
	Type        int             `bson:"type" json:"type"`         // 0: 私密文章, 1: 公开文章, 2: 简历, 3: 管理员介绍
	Operator    int             `bson:"operator" json:"operator"` // 0: 草稿, 1: 发布
	Isdelete    bool            `bson:"isdelete" json:"isdelete"`
}
type Meta struct {
	Watch int `bson:"watch" json:"watch"`
	Like  int `bson:"like" json:"like"`
	Num   int `bson:"num" json:"num"`
}

/*
// 增加用户
func InsertUser(user *User) bool {
	flag := Insert("user", user)
	return flag
}
*/
func InsertArticle(article *Article) bool {
	flag := Insert("article", article)
	return flag
}

func FindArticleById(id string) (*Article, bool) {
	article := new(Article)
	filter := bson.M{}
	filter["_id"] = bson.ObjectIdHex(id)
	filter["isdelete"] = false
	flag := Find("article", filter, &article)
	return article, flag
}
