package objects

import (
	"time"

	"github.com/gorilla/websocket"
)

type LogData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
	Time       time.Time
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
	Time      time.Time
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
	Connection   *websocket.Conn
}

type Chat struct{}

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

type Error struct {
	StatusCode   int
	ErrorMessage string
}
