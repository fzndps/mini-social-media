package domain

type CommentWitUser struct {
	Id        int
	Content   string
	UserId    int
	PostId    int
	CreatedAt string
	User      UserCommentInfo
}

type UserCommentInfo struct {
	Id       int
	Username string
}

type PostWithComments struct {
	Id        int
	UserId    int
	Username  string
	Content   string
	ImageURL  string
	CreatedAt string
	Comments  []CommentWitUser
}
