package dao

import "time"

//Post 文章内容
type Post struct {
	ID       string    `bson:"_id"`
	Title    string    `bson:"title"`
	URL      string    `bson:"url"`
	PostDate time.Time `bson:"postDate"`
	Desc     string    `bson:"desc"`
}
