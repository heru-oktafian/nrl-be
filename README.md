# NRL-BE

Backend API for Nurul Deere portfolio website.

## Tech Stack

- **Go** 1.21+
- **Fiber** v2 (Web Framework)
- **PostgreSQL** (Database)
- **pgx** v5 (PostgreSQL Driver)

## Architecture

Clean Architecture pattern with the following structure:

```
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/              # Configuration and database setup
│   ├── delivery/
│   │   └── http/            # HTTP handlers and routes
│   ├── domain/
│   │   ├── entity/         # Domain entities
│   │   └── repository/     # Repository interfaces
│   └── infrastructure/      # External services
├── pkg/
│   └── response/           # Response utilities
└── docs/                   # Documentation
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 14+

### Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your database settings
3. Create the database:
   ```bash
   createdb nrl_be
   ```
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

### Environment Variables

```env
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=nrl_be
APP_PORT=3001
```

## API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/portfolio` | Get all portfolio data |
| GET | `/api/v1/profile` | Get profile |
| GET | `/api/v1/experiences` | Get experiences |
| GET | `/api/v1/skills` | Get skills |
| GET | `/api/v1/projects` | Get projects |
| GET | `/api/v1/social-links` | Get social links |
| GET | `/api/v1/tools` | Get tools |
| POST | `/api/v1/contact` | Submit contact message |

### Contact Request Example

```json
POST /api/v1/contact
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "subject": "Hello",
  "message": "I would like to get in touch..."
}
```

## Database Schema

The application automatically runs migrations on startup. Tables created:

- `profiles` - Profile information
- `experiences` - Work experiences
- `skills` - Skills data
- `projects` - Project portfolio
- `social_links` - Social media links
- `tools` - Tools used
- `contact_messages` - Contact form submissions

## Development

### Run locally

```bash
go run cmd/api/main.go
```

### Build

```bash
go build -o bin/nrl-be cmd/api/main.go
```

## License

MIT