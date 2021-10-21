package models

type PostModel struct {
	ID           int
	Title        string
	Category     string
	CategoryRank int `json:"-" db:"category_rank"`
}

var posts []PostModel

func (p PostModel) Reset() {
	posts = nil
}
func AddPost(pm PostModel) {
	posts = append(posts, pm)
}
func (p PostModel) GetPosts() []PostModel {
	return posts
}
