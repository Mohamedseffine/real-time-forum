package objects

import "time"

type Post struct {
	Username   string
	UserId     int
	Ncomments  int
	Content    string
	Title      string
	Categories []string
	PostId     int
	Reacts     []PostReaction
	Comments   []Comment
}

type PostReaction struct {
	ReactId int
	UserId  int
	PostId  int
	Type    string
}

type CommentReact struct {
	ReactId   int
	UserId    int
	Username  string
	PostId    int
	Type      string
	CommentId int
}

type Comment struct {
	UserId    int
	PostId    int
	Username  string
	CommentId int
	Content   string
	Reacts    []CommentReact
}

type User struct {
	Id           int
	FamilyName   string
	Name         string
	Username     string
	BirthDate    time.Time
	Gender       string
	Email        string
	CreatedPosts []Post
	LikedPosts   []Post
	State        bool
}

type Chat struct {
}

type Message struct {
	MessageId  int
	UserId     int
	ChatId     int
	Content    string
	Date       time.Time
	Username   string
	Reciever   string
	RecieverId int
}
