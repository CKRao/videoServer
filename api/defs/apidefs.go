package defs

//用户结构体
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//视频信息结构体
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

//评论结构体
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

//会话结构体
type SimpleSession struct {
	UserName string //login name
	TTL      int64
}
