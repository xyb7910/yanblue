package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp is the struct for sign up
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

// ParamLogin is the struct for login
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData is the struct for vote
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof = 1 0 -1"`
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page" form:"page" example:"1"`
	Size        int64  `json:"size" form:"size" example:"10"`
	Order       string `json:"order" form:"order" example:"score"`
}
