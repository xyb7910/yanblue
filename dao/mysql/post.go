package mysql

import "yanblue/models"

// CreatePost create post
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
                 post_id,title,content,author_id,community_id) 
	values (?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostByID get post by id
func GetPostByID(pid int64) (post *models.Post, err error) {
	post = &models.Post{}
	sqlStr := `select 
    post_id,title,content,author_id,community_id,create_time   
	from post where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
    post_id,title,content,author_id,community_id,create_time   
	from post order by create_time desc limit ?,?`

	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
