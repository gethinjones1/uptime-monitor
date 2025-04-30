# 🟢 Uptime Monitor

**Uptime Monitor** is a full-stack, containerized platform for monitoring the availability and health of services and endpoints. It includes a real-time dashboard, a RESTful API, background workers, and a CLI — all orchestrated using Docker Compose.

---

## 🔧 Components

### 📊 Dashboard (`dashboard/`)
- Frontend app built with **React** and **TanStack Query**
- Displays real-time service status
- Runs independently or inside Docker

### 🌐 API (`api/`)
- Written in **Go**
- Serves RESTful endpoints
- Includes database migrations via the `migrate` service
- Hot reload enabled via [Air](https://github.com/cosmtrek/air)

### 🛠 CLI (`cmd/`)
- Go-based command-line tool for managing and interacting with the monitoring system
- Useful for tasks like seeding checks, running diagnostics, or manual invocations

### ⚙️ Worker (`worker/`)
- Background process written in Go
- Periodically pings endpoints, records status to the database
- Designed to run continuously as a long-lived service

---

## 🐳 Dockerized Setup

All components are containerized and orchestrated with Docker Compose.

### ▶️ To Start

```bash
docker-compose up --build
cd dashboard
npm run dev
