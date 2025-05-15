-- Users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    gender VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    creation_date DATETIME NOT NULL
);
-- Posts table
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    creator_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    username VARCHAR(50) NOT NULL,
    creation_time DATETIME NOT NULL,
    content TEXT,
    FOREIGN KEY (creator_id) REFERENCES users (id) ON DELETE CASCADE
);
-- Categories table
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    category VARCHAR(50) UNIQUE NOT NULL
);
-- Insert default categories
INSERT
    OR IGNORE INTO categories (category)
VALUES ('Sport'),
    ('Music'),
    ('Movies'),
    ('Science'),
    ('Politics'),
    ('Culture'),
    ('Technology');
-- Comments table
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    creator_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    username VARCHAR(50) NOT NULL,
    creation_date DATETIME NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE
);
-- Messages table
CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    mtype VARCHAR(20) NOT NULL,
    recieved_at DATETIME NOT NULL,
    sender_username VARCHAR(50) NOT NULL,
    reciever_username VARCHAR(50) NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE
);
-- Sessions table
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    token VARCHAR(500) NOT NULL,
    created_at DATETIME NOT NULL,
    expires_at DATETIME NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS post_categories (
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS chats (
    chat_id INTEGER PRIMARY KEY,
    unread_messages INTEGER NOT NULL,
    st_part VARCHAR(50) NOT NULL,
    nd_part VARCHAR(50) NOT NULL
);
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL