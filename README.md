# hack4good

## Backend Setup

### Prerequisites
- Install Go
- Install Docker + Docker Compose

### Steps
In backend directory:
1. Start PostgreSQL database:
`docker compose up -d`

2. Create `.env` file with:
```
DATABASE_URL=postgresql://devuser:devpassword@host.docker.internal:5432/hack4good
JWT_SECRET=change-this
PORT=8080
```

3. Install dependencies:
`go mod tidy`

4. Run server:
`go run ./cmd/main.go`
