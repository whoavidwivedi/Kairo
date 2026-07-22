# Kairo

> **Kairo is an AI-powered CRM built to help businesses capture, organize, and manage leads through intelligent automation.**

Kairo combines a modern web application with a scalable Go backend to provide a fast, reliable foundation for customer relationship management and future AI-powered workflows.

---

## 🚀 Tech Stack

### Frontend

- Next.js
- React
- TypeScript
- Tailwind CSS
- Bun

### Backend

- Go
- Chi Router

### Planned

- PostgreSQL
- Docker
- OpenAI API
- JWT Authentication
- Redis
- GitHub Actions

---

## 📂 Project Structure

```text
Kairo/
├── apps/
│   ├── api/          # Go backend
│   └── web/          # Next.js frontend
├── .gitignore
├── bun.lock
├── package.json
└── README.md
```

---

## 🏗️ Backend Architecture

Kairo follows a layered architecture to keep the codebase maintainable and easy to scale.

```text
Client
   │
   ▼
HTTP Route
   │
   ▼
Service Layer
   │
   ▼
Repository Layer
   │
   ▼
Storage
(Currently In-Memory)
```

### Layers

#### Models

Contains the application's data structures.

#### Services

Contains business logic and coordinates application behavior.

#### Repository

Handles data persistence and abstracts storage implementation.

#### Routes

Receives HTTP requests and returns HTTP responses.

---

## ✨ Current Features

- Health check endpoint
- Create Lead API
- List Leads API
- Layered backend architecture
- Dependency Injection
- In-memory lead storage
- RESTful API design

---

## 🗺️ Roadmap

### ✅ Backend Foundation

- [x] Initialize Go backend
- [x] Configure Chi router
- [x] Create project architecture
- [x] Implement LeadService
- [x] Implement LeadRepository
- [x] Add Health endpoint
- [x] Add Create Lead endpoint
- [x] Add List Leads endpoint

### 🚧 Database Integration

- [ ] PostgreSQL integration
- [ ] SQL migrations
- [ ] Persistent storage
- [ ] Repository refactor

### 🔐 Authentication

- [ ] JWT Authentication
- [ ] User accounts
- [ ] Protected routes

### 🤖 AI Features

- [ ] AI lead insights
- [ ] Automated follow-ups
- [ ] CRM assistant
- [ ] Workflow automation

### 🚀 Production

- [ ] Docker
- [ ] CI/CD
- [ ] Monitoring
- [ ] Deployment

---

## 🛠️ Getting Started

### Prerequisites

- Bun
- Go 1.24+

---

### Clone the repository

```bash
git clone https://github.com/whoavidwivedi/Kairo.git
cd Kairo
```

---

### Install frontend dependencies

```bash
bun install
```

---

### Run the frontend

```bash
cd apps/web
bun dev
```

The frontend will be available at:

```text
http://localhost:3000
```

---

### Run the backend

```bash
cd apps/api
go run ./cmd/server
```

The backend will be available at:

```text
http://localhost:8080
```

---

## 📡 API Endpoints

### Health Check

```http
GET /health
```

Response

```json
{
  "status": "ok"
}
```

---

### Create Lead

```http
POST /leads
```

Example Request

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "1234567890"
}
```

---

### List Leads

```http
GET /leads
```

---

## 🧪 Current Status

The backend currently stores data in memory.

Database persistence using PostgreSQL is the next major milestone.

---