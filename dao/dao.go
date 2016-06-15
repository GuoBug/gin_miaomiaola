package dao

import (
	"time"

	"github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Add 添加话题
func (p *Post) Add(db *mgo.Database, log log15.Logger) error {
	c := db.C(CollectionName)
	p.PostDate = time.Now()
	err := c.Insert(p)
	if err != nil {
		log.Error("新增出错", "Params", p, "error", err)
	}
	return err
}

//GetAllTopic 获取全部话题
func GetAllTopic(db *mgo.Database, log log15.Logger) (*[]Post, error) {
	p := new([]Post)
	cl := db.C(CollectionName)
	q := cl.Find(bson.M{"_id": bson.M{"$ne": "555ad4e8e64786646a000018"}}).Sort("-createTime")
	err := q.All(p)
	if err != nil {
		log.Error("获取全部失败")
	}
	return p, err
}

//GetTopic 获取全部话题
func GetTopic(url string, db *mgo.Database, log log15.Logger) (*Post, error) {
	p := new(Post)
	cl := db.C(CollectionName)
	q := cl.Find(bson.M{"url": url})
	err := q.One(p)
	if err != nil {
		log.Error("获取单个失败")
	}
	return p, err
}
