package models

type PostModel struct {
	ID    int
	Title string
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
