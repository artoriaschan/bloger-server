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
	Creater    bson.ObjectId `bson:"creater" json:"creater"`
	Isdelete   bool          `bson:"isdelete" json:"isdelete"`
}
type OutputCategory struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Catename   string        `bson:"catename" json:"catename"`
	Createtime int64         `bson:"createtime" json:"createtime"`
	Updatetime int64         `bson:"upadtetime" json:"upadtetime"`
	Creater    bson.ObjectId `bson:"creater" json:"creater"`
}

// 插入分类
func InsertCategory(category *Category) bool {
	flag := Insert("category", category)
	return flag
}

// 获取分类列表
func GetCategories(filter bson.M, skip, limit int) (*[]OutputCategory, int, bool) {
	cateList := new([]OutputCategory)
	field := bson.M{
		"_id":        1,
		"catename":   1,
		"createtime": 1,
		"upadtetime": 1,
		"creater":    1,
	}
	filter["isdelete"] = false
	countNum, flag := FindAll("category", filter, field, cateList, skip, limit)
	return cateList, countNum, flag
}

func UpdateCategory(id string, catename string) bool {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	data := bson.M{"$set": bson.M{"catename": catename}}
	flag := Update("category", selector, data)
	return flag
}

func DeleteCategory(id string) bool {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	data := bson.M{"$set": bson.M{"isdelete": true}}
	flag := Update("category", selector, data)
	return flag
}
