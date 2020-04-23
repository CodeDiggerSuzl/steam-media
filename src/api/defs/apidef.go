package defs

// UserCredential Response definition
type UserCredential struct {
	UserName string `json:"user_name"`
	PassWard string `json:"password"`
}

// VideoInfo represent for video
type VideoInfo struct {
	ID                string
	AuthorID          string
	Name              string
	DisplayCreateTime string
}
