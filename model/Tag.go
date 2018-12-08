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
	Tagname    string        `bson:"tagname" json:"catename"`
	Color      string        `bson:"color" json:"color"`
	Createtime int64         `bson:"createtime" json:"createtime"`
	Updatetime int64         `bson:"upadtetime" json:"upadtetime"`
	Creater    User          `bson:"creater" json:"creater"`
	Isdelete   bool          `bson:"isdelete" json:"isdelete"`
}
