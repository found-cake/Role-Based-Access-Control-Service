# Role-Based-Access-Control-Service (Go)

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
