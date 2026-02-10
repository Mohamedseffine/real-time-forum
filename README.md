# Real‑Time Forum — Go + WebSockets (Chat & Notifications)

A full‑stack **forum application** with **real‑time notifications** and **live chat** using **WebSockets**.
Users can sign up / log in, create posts and comments, and receive updates instantly without refreshing.

---

## Features
- ✅ Authentication: **signup / login / logout**
- ✅ Forum: create posts, view posts, create comments, list comments
- ✅ Real‑time **chat** (WebSockets)
- ✅ Real‑time **notifications** (WebSockets)
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
```

---

## How It Works

### REST API (forum + auth)
- Auth endpoints handle **signup/login/logout**
- Forum endpoints handle:
  - Create a post
  - Get post(s)
  - Create comment(s)
  - Get comments

These handlers live under: `backend/handlers/`

### WebSockets (chat + notifications)
WebSockets are used for events that should arrive instantly:
- New chat messages
- New notifications (e.g., someone commented on your post)

The WebSocket upgrade/handshake logic is typically managed by:
- `backend/handlers/upgrade_protocol.go`
- `backend/handlers/chat_handler.go`

---

## Run Locally

### 1) Requirements
- Go installed (matching your `go.mod`)
- A browser for the frontend

### 2) Start the project
From the project root:

```bash
go run .
```


```bash
go run main.go
```




## Database
- SQLite database file: `backend/database/database.db`
- SQL schema: `backend/database/schema.sql`

If you need to recreate the DB, delete `database.db` and run your app again (if your code auto‑creates tables),
or apply the schema manually (depending on your implementation).

---

## Middleware
Located in `backend/middleware/`

- **auth.go**  
  Protects routes that require a logged-in user (usually by checking a session/cookie/token).

- **rate_limiter.go**  
  Limits request frequency to protect the API from abuse/spam.

---

## Backend Packages (Quick Map)

- `backend/api/api.go`  
  API wiring / route registration (commonly).

- `backend/handlers/`  
  HTTP handlers (REST + WS).

- `backend/models/`  
  Data models + DB access helpers.

- `backend/helpers/`  
  Reusable helper logic (e.g., login utilities).

- `backend/objects/`  
  Shared DTOs / structs for passing data between layers.

---

## WebSocket Events (Suggested Convention)

Your exact message format depends on your implementation, but a common pattern is JSON messages like:

```json
{
  "type": "chat_message",
  "payload": {
    "from": "userA",
    "to": "userB",
    "message": "Hello!",
    "sentAt": "2026-02-10T12:00:00Z"
  }
}
```

And for notifications:

```json
{
  "type": "notification",
  "payload": {
    "kind": "new_comment",
    "postId": 42,
    "message": "Someone commented on your post",
    "createdAt": "2026-02-10T12:00:00Z"
  }
}
```

---


