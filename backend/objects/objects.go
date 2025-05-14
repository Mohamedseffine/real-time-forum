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
	ID           int    `json:"id"`
	UserId       int    `json:"creator_id"`
	Username     string `json:"username"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreationTime string `json:"creation_time"`
	Categories   []int  `json:"categories"`
	Categorie    string `json:"categorie"`
	Ncomments    int
	Reacts       []PostReaction
	Comments     []Comment
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
	CommentId int       `json:"comment_id"`
	UserId    int       `json:"user_id"`
	PostId    int       `json:"post_id"`
	Username  string    `json:"username"`
	Time      time.Time `json:"time"`
	Content   string    `json:"content"`
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

type Chat struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	MessageId  int
	UserId     int       `json:"user_id"`
	ChatId     int       `json:"chat_id"`
	Content    string    `json:"message"`
	Dare       time.Time `json:"time"`
	Type       string    `json:"read_unread"`
	Username   string    `json:"username"`
	Reciever   string    `json:"reciever"`
	RecieverId int       `json:"reciever_id"`
}

type Error struct {
	StatusCode   int
	ErrorMessage string
}

type WsData struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Users   []Infos `json:"users"`
}

type Infos struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IsActive int    `json:"active"`
}

var Users = make(map[int]*websocket.Conn, 24)
