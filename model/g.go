package model
import (
	"gopkg.in/mgo.v2"
	"log"
)

var GlobalMgoSession *mgo.Session
var dbName string = "bloger"

func init(){
	db := ConnectToDB()
	SetDB(db)
}

func SetDB(session *mgo.Session){
	GlobalMgoSession = session
}

func ConnectToDB() *mgo.Session{
	GlobalMgoSession, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	GlobalMgoSession.SetPoolLimit(10)

	return GlobalMgoSession
}
// 插入
func Insert(collectionName string, documents interface{}) bool{
	session := GlobalMgoSession.Clone()
	defer session.Close()

	collection := session.DB(dbName).C(collectionName)

	err := collection.Insert(documents)

	if err != nil {
		log.Println( err.Error())
		return false
	}
	return true
}
// 查找
func Find(collectionName string, M interface{}, result interface{}) bool{
	session := GlobalMgoSession.Clone()
	defer session.Close()

	collection := session.DB(dbName).C(collectionName)

	err := collection.Find(M).One(result)
	if err != nil {
		log.Println( err.Error())
		return false
	}
	return true
}
// 删除
func Remove(collectionName string){
	session := GlobalMgoSession.Clone()
	defer session.Close()

	//collection := session.DB(dbName).C(collectionName)
}
// 更新
func Update(collectionName string){
	session := GlobalMgoSession.Clone()
	defer session.Close()

	//collection := session.DB(dbName).C(collectionName)
}
