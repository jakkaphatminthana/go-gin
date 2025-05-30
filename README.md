# Go-Gin API

A simple RESTful API project built with [Gin](https://github.com/gin-gonic/gin), designed as a playground for learning modern Go web development concepts such as routing, middleware, database integration, and OAuth2 login.

## 🚀 Features

- RESTful API with CRUD support
- PostgreSQL integration using GORM
- Google OAuth2 login
- Environment-based configuration (via `viper`)
- Clean project structure with separation of concerns
- Elastic Beanstalk-ready deployment

## 🧱 Tech Stack

- Go 1.21+
- Gin framework
- PostgreSQL
- GORM ORM
- Viper (config)
- OAuth2 (Google login)
- Elastic Beanstalk (deployment)

## 📂 Project Structure

```text
go-gin/
├── config/           # Configuration loading (viper)
├── database/         # Database initialization and migrations
├── entities/         # GORM models
├── pkg/
│   ├── auth/         # OAuth2 logic
│   └── task/         # Task CRUD logic
├── server/           # Gin setup and routes
└── main.go           # Entry point
```

## ⚙️ Configuration
Create a `.env` file to populate environment variables:

```env
APP_PORT=8080 #5000 production

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=tasks
SSL_MODE=disable

GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_REDIRECT_URL=your-google-redirect-url

JWT_SALT_KEY=
FE_ORIGINAL_URL=
```

Then set up config.yaml (already in repo)

## 🛠 Setup

#### 1. Install dependencies
``` go mod tidy ```

#### 2. Run migration
``` go run database/migrations/migration.go ```

#### 3. Start server
``` go run main.go ```

## ☁️ Deployment to AWS Elastic Beanstalk
```
eb init -p go go-gin-dev
eb create go-gin-dev
eb deploy
```

Connect command: ssh -i ~/.ssh/aws-eb -N -L 5432:<END_POINT>:5432 ec2-user@<EC2_IPV4>
