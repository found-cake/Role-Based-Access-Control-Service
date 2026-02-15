# Role-Based-Access-Control-Service (Go)

Node.js/TypeScript 구현을 Go로 전환한 RBAC 인증 서비스입니다.

## Stack

- Go 1.26+
- Echo (`github.com/labstack/echo/v4`)
- GORM (`gorm.io/gorm`)
- PostgreSQL 15
- JWT (`github.com/golang-jwt/jwt/v5`)
- Request validator (`github.com/go-playground/validator/v10`)
- Password hashing (`bcrypt`)

## Quick Start

1. PostgreSQL 실행

```bash
docker compose up -d
```

2. 환경 변수 설정

```bash
cp .env.example .env
```

3. 실행

```bash
make run
```

## Environment Variables

- `PORT` (default: `3000`)
- `DB_USER` (default: `user`)
- `DB_PASSWORD` (default: `password`)
- `DB_HOST` (default: `localhost`)
- `DB_PORT` (default: `5433`)
- `DB_NAME` (default: `rbac_db`)
- `JWT_SECRET` (default: `secret`)

## Package Layout

- `main.go`: app entrypoint
- `handlers`: HTTP handler layer
- `service`: business logic layer
- `db`: connection + repository layer
- `dto`: request/response DTOs
- `models`: DB models
- `middleware`: auth middleware
- `validation`: Echo validator adapter
- `pkg/auth`: JWT helpers
- `pkg/httpx`: HTTP response helpers

## API

- `POST /auth/register`
- `POST /auth/login`
- `GET /auth/me` (Authorization: Bearer <token>)
- `GET /health`

정적 파일은 `public/`에서 서빙됩니다.
