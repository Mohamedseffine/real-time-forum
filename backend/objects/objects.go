package objects

import (
	"time"

	"github.com/gorilla/websocket"
)

type LogData struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	FamilyName   string    `json:"lastname"`
	Name         string    `json:"firstname"`
	Gender       string    `json:"gender"`
	CreationDate time.Time `json:"creation_date"`
	LogType      string    `json:"type"`
}

type Post struct {
	Username   string `json:"username"`
	UserId     int    `json:"id"`
	Content    string `json:"content"`
	Title      string `json:"title"`
	Categories []int  `json:"categories"`
	Ncomments  int
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
	UserId    int    `json:"id"`
	PostId    int    `json:"postid"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CommentId int
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
	SessionToken string
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
