# Real-Time Forum — Go + WebSockets (Chat & Notifications)

A full-stack **forum application** with **real-time notifications** and **live chat** using **WebSockets**.
Users can sign up / log in, create posts and comments, and receive updates instantly without refreshing.

---

## Features
- ✅ Authentication: **signup / login / logout**
- ✅ Forum: create posts, view posts, create comments, list comments
- ✅ Real-time **chat** (WebSockets)
- ✅ Real-time **notifications** (WebSockets)
- ✅ Middleware:
  - Auth protection for private routes
  - Rate limiting
- ✅ SQLite database with schema file

---

## Tech Stack
- **Backend:** Go (net/http), WebSockets
- **Database:** SQLite
- **Frontend:** HTML/CSS/JavaScript (in `frontend/`)
- **Protocols:** REST for standard operations + WebSockets for live events


---

## Project Structure

```txt
.
├── backend/
│   ├── api/
│   │   └── api.go
│   ├── handlers/
│   │   ├── chat_handler.go
│   │   ├── create_comments.go
│   │   ├── create_posts.go
│   │   ├── error_handler.go
│   │   ├── get_comments.go
│   │   ├── get_post.go
│   │   ├── home_handler.go
│   │   ├── login.go
│   │   ├── logout.go
│   │   ├── ressourcesHandler.go
│   │   ├── signup.go
│   │   └── upgrade_protocol.go
│   ├── helpers/
│   │   └── login_helper.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── rate_limiter.go
│   ├── models/
│   │   ├── chat.go
│   │   ├── comments.go
│   │   ├── database.go
│   │   ├── posts.go
│   │   └── users.go
│   ├── objects/
│   │   └── objects.go
│   └── database/
│       ├── database.db
│       └── schema.sql
├── frontend/
├── main.go
├── go.mod
├── go.sum
├── .gitignore
├── push.sh
├── script.sh
└── ho.txt
