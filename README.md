# Go Login API

RESTful API สำหรับระบบ Login/Register ด้วย JWT Token

## โครงสร้างโปรเจค

```
GO_LOGIN/
├── cmd/                    # Entry point of application
│   └── main.go            # Main application file
├── internal/              # Private application code
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # HTTP middlewares
│   ├── models/            # Data models and database operations
│   └── routes/            # Route definitions
├── pkg/                   # Public packages
│   └── auth/             # JWT authentication utilities
├── scripts/              # Utility scripts
│   └── generate_key.go   # JWT key generation
├── .env                  # Environment variables
├── docker-compose.yml    # Docker services configuration
├── Dockerfile            # Container build instructions
└── go.mod               # Go module definition
```

## การติดตั้ง

1. Clone repository:
```bash
git clone https://github.com/WUT-WORAWIT/go_login.git
cd go_login
```

2. ติดตั้ง dependencies:
```bash
go mod download
```

3. สร้าง JWT secret key:
```bash
go run scripts/generate_key.go
```

4. ตั้งค่า environment variables ใน `.env`

## การรัน

### Local Development
```bash
go run cmd/main.go
```

### Docker
```bash
docker-compose up --build
```

## API Endpoints

- POST /api/login - เข้าสู่ระบบ
- POST /api/users - สมัครสมาชิก
- GET /api/users - ดูข้อมูลผู้ใช้ทั้งหมด (ต้องมี JWT Token)
- GET /api/users/:id - ดูข้อมูลผู้ใช้รายบุคคล (ต้องมี JWT Token)

## เทคโนโลยีที่ใช้

- Go 1.24
- Gin Web Framework
- GORM
- PostgreSQL
- JWT Authentication
- Docker
