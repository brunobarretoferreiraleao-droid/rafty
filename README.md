# 🚀 Rafty

> **A modern social network MVP built for speed, interaction, and clean architecture.**

**Rafty** is a full-stack social media application designed initially as a **minimum viable product (MVP)** with a strong foundation for future scaling.

Built with:

- **Next.js**
- **TypeScript**
- **Tailwind CSS**
- **Go (Fiber)**
- **PostgreSQL**
- **JWT Authentication**

It focuses on the core mechanics of a real social platform: **users, posts, follows, friendships, conversations, and media support**.

---

## ✨ Vision

Rafty is meant to be:
 
- **Fast**
- **Modern**
- **Modular**
- **Scalable**

The goal is to create a social platform that feels lightweight and expressive, while also serving as a serious **full-stack engineering project**.

---

# 📸 Core MVP Features

## 👤 Authentication
- User registration
- Secure login
- Password hashing with **bcrypt**
- JWT-based authentication
- Protected user session routes

## 📝 Posts
- Create text posts
- Visibility control:
  - `public`
  - `friends`
  - `private`
- Support for future media expansion

## 🖼️ Media Attachments
Posts and messages are structured to support:

- Images
- Audio
- Videos
- GIFs
- Generic files

## 🔎 User Discovery
- Search users by username
- Profile-oriented architecture

## 🤝 Friendship System
- Send friend requests
- Accept / reject requests
- Bidirectional friendship model

## 👀 Follow System
- Follow other users
- Designed for public-post notification behavior
- Inspired by platforms like Twitter / X

## 💬 Private Messaging
- One-to-one conversations
- Direct messages
- Message editing support
- Message attachments
- Conversation ordering by latest activity

---

# 🏗️ Tech Stack

## Frontend
- **Next.js**
- **TypeScript**
- **Tailwind CSS**
- **ESLint**
- **Prettier**

## Backend
- **Go**
- **Fiber**
- **JWT**
- **bcrypt**

## Database
- **PostgreSQL**
- **UUID**
- **CITEXT**
- **PL/pgSQL triggers & validation**

---

# 📂 Project Structure

```bash
rafty/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── config/
│   │   ├── controllers/
│   │   ├── middlewares/
│   │   ├── models/
│   │   ├── repositories/
│   │   ├── routes/
│   │   ├── services/
│   │   └── utils/
│   └── go.mod
│
├── database/
│   └── schema.sql
│
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   ├── components/
│   │   ├── lib/
│   │   ├── hooks/
│   │   ├── types/
│   │   └── constants/
│   └── package.json
│
└── README.md

---

⚡ Quick Start (How to run that app)
1) Clone the repository
git clone https://github.com/brunobarretoferreiraleao-droid/rafty.git
cd rafty

2) Setup the database

Create a PostgreSQL database and run:

psql -U postgres -d rafty -f database/schema.sql
3) Configure backend environment

Create:

backend/.env

Example:

PORT=8080
DATABASE_URL=postgres://postgres:your_password@localhost:5432/rafty?sslmode=disable
JWT_SECRET=change_this_secret
JWT_EXPIRES_HOURS=24

4) Run the backend
cd backend
go mod tidy
go run cmd/main.go

Expected API base URL:

http://localhost:8080

5) Run the frontend
cd frontend
npm install
npm run dev

Open:

http://localhost:3000

---

🧠 Architecture Highlights

Rafty was designed with a layered backend architecture:

Controllers → HTTP request handling
Services → business logic
Repositories → database access
Middlewares → auth / request validation
Utils → shared helpers like JWT and password hashing

This structure keeps the backend clean, testable, and maintainable.

🗃️ Database Design

The PostgreSQL schema already includes support for:

users
posts
post_attachments
follows
friends
friend_requests
conversations
messages
message_attachments

It also includes:

triggers
custom validation functions
pair uniqueness rules
automatic updated_at handling
conversation activity tracking
🔐 Authentication Flow

Rafty uses a standard JWT authentication flow:

User registers or logs in
Backend validates credentials
JWT token is generated
Frontend stores token
Protected routes use Authorization: Bearer <token>

This enables a clean separation between frontend and backend while keeping the MVP implementation practical and scalable.