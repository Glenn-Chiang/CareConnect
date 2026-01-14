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
DATABASE_URL=postgres://postgres:postgres@localhost:5432/hack4good?sslmode=disable
JWT_SECRET=change-this
PORT=8080
GIN_MODE=debug
```

3. Install dependencies:
`go mod tidy`

4. Run server:
`go run ./cmd/main.go`
