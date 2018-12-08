package model

import "gopkg.in/mgo.v2/bson"

/*
{
    id ObjectId,    *
    catename string,    *
    createtime timestemp,   *
    upadtetime timestemp,   *
    creater User{}, *
    isdelete bool   *
}
*/
type Category struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Catename   string        `bson:"catename" json:"catename"`
	Createtime int64         `bson:"createtime" json:"createtime"`
	Updatetime int64         `bson:"upadtetime" json:"upadtetime"`
	Creater    User          `bson:"creater" json:"creater"`
	Isdelete   bool          `bson:"isdelete" json:"isdelete"`
}
