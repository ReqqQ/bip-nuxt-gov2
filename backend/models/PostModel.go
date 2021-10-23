package models

import "time"

type PostModel struct {
	ID           int
	Title        string
	Category     string
	CategoryRank int `json:"-" db:"category_rank"`
	Timestamp    int `json:"-" db:"created_at"`
	CreatedAt    string
}

var posts []PostModel

func (pm PostModel) Reset() {
	posts = nil
}
func (pm *PostModel) transformDate() {
	pm.CreatedAt = time.Unix(int64(pm.Timestamp), 0).Format("02-January-2006 15:04:05")
}
func AddPost(pm PostModel) {
	pm.transformDate()
	posts = append(posts, pm)
}
func (pm PostModel) GetPosts() []PostModel {
	return posts
}
