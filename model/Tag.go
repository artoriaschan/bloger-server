package model

import "gopkg.in/mgo.v2/bson"

/*
{
    id ObjectId,    *
    tagname string,    *
	color string,	*
    createtime timestemp,   *
    upadtetime timestemp,   *
    creater User{}, *
    isdelete bool   *
}
*/
type Tag struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Tagname    string        `bson:"tagname" json:"tagname"`
	Color      string        `bson:"color" json:"color"`
	Createtime int64         `bson:"createtime" json:"createtime"`
	Updatetime int64         `bson:"upadtetime" json:"upadtetime"`
	Creater    bson.ObjectId `bson:"creater" json:"creater"`
	Isdelete   bool          `bson:"isdelete" json:"isdelete"`
}

type OutputTag struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Tagname    string        `bson:"tagname" json:"tagname"`
	Color      string        `bson:"color" json:"color"`
	Createtime int64         `bson:"createtime" json:"createtime"`
	Updatetime int64         `bson:"upadtetime" json:"upadtetime"`
	Creater    bson.ObjectId `bson:"creater" json:"creater"`
}

// 插入标签
func InsertTag(tag *Tag) bool {
	flag := Insert("tag", tag)
	return flag
}

// 获取分类列表
func GetTags(filter bson.M, skip, limit int) (*[]OutputTag, int, bool) {
	cateList := new([]OutputTag)
	field := bson.M{
		"_id":        1,
		"tagname":    1,
		"color":      1,
		"createtime": 1,
		"upadtetime": 1,
		"creater":    1,
	}
	filter["isdelete"] = false
	countNum, flag := FindAll("tag", filter, field, cateList, skip, limit)
	return cateList, countNum, flag
}
