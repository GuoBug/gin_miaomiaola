package dao

import (
	"github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetAllTopic 获取全部话题
func GetAllTopic(db *mgo.Database, log log15.Logger) (*[]Post, error) {
	p := new([]Post)
	cl := db.C(CollectionName)
	q := cl.Find(bson.M{}).Sort("-createTime")
	err := q.All(p)
	if err != nil {
		log.Error("添加用户失败")
	}
	return p, err
}

//GetTopic 获取全部话题
func GetTopic(db *mgo.Database, log log15.Logger) (*Post, error) {
	p := new(Post)
	cl := db.C(CollectionName)
	q := cl.Find(bson.M{})
	err := q.One(p)
	if err != nil {
		log.Error("添加用户失败")
	}
	return p, err
}
