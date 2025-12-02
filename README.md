# Project_5_E-Commerce

Folder Structure -----


ecommerce-backend/
│
├── cmd/                      # Entry points of your application
│   └── server/               # Main API server
│       └── main.go
│
├── internal/                 # Application internal logic (private to your app)
│   ├── config/               # Configuration (DB, environment variables, JWT, etc.)
│   │   └── config.go
│   │
│   ├── models/               # Database models / structs
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── order.go
│   │   └── ...
│   │
│   ├── repositories/         # Database queries (GORM or raw SQL)
│   │   ├── user_repo.go
│   │   ├── product_repo.go
│   │   └── order_repo.go
│   │
│   ├── services/             # Business logic layer
│   │   ├── user_service.go
│   │   ├── product_service.go
│   │   ├── order_service.go
│   │   └── payment_service.go
│   │
│   ├── handlers/             # HTTP handlers (controllers)
│   │   ├── user_handler.go
│   │   ├── product_handler.go
│   │   └── order_handler.go
│   │
│   ├── middleware/           # Auth, logging, rate limiting, CORS, etc.
│   │   ├── auth.go
│   │   └── logger.go
│   │
│   ├── utils/                # Helper functions, constants
│   │   ├── validator.go
│   │   ├── crypto.go
│   │   └── email.go
│   │
│   ├── routes/               # API routes
│   │   └── routes.go
│   │
│   ├── cache/                # Redis or other caching logic
│   │   └── cache.go
│   │
│   ├── notifications/        # Email, SMS, push notification services
│   │   └── email_service.go
│   │
│   └── storage/              # File storage (S3 or local)
│       └── s3.go
│
├── pkg/                      # External libraries/utilities that can be shared
│   └── logger/               # Custom logging package
│
├── scripts/                  # DB migrations, seeders, deployment scripts
│   └── migrate.go
│
├── migrations/               # DB migration files (if using sql-migrate or goose)
│
├── docs/                     # API documentation (Swagger / Postman collections)
│
├── go.mod
├── go.sum
├── .env                      # Environment variables
└── README.md