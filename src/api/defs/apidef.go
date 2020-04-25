package defs

// UserCredential Response definition
type UserCredential struct {
	UserName string `json:"user_name"`
	PassWard string `json:"password"`
}

// VideoInfo represent for video
type VideoInfo struct {
	ID                string
	AuthorID          int
	Name              string
	DisplayCreateTime string
}

// Comment for video comment
type Comment struct {
	ID      string
	Author  string
	VideoID string
	Content string
}

// Session struct
type Session struct {
	UserName string
	TTL      int64 // time to live, use to check if the login info is expired or not
}