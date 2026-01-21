# ARGolang

[![Go Version](https://img.shields.io/github/go-mod/go-version/DanielP41/API-REST-Golang)](https://go.dev/)

Robust REST API base project developed in Go, utilizing PostgreSQL for data persistence. Designed with a clean and modular architecture.

## System Architecture

![Project Architecture](docs/assets/architecture.png)

---

## Features

- Full CRUD for tasks.
- Input data validation.
- Robust PostgreSQL connection via sqlx.
- Environment variable management with .env.
- CI/CD pipeline configured with GitHub Actions.
- Standardized folder structure for Go projects.

## Technologies

- **Language:** [Go](https://go.dev/) (v1.21+)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Core Libraries:**
  - `sqlx`: Extension to database/sql for easier querying.
  - `godotenv`: Environment variable loading.
  - `lib/pq`: PostgreSQL driver for Go.

## Project Structure

```text
.
├── cmd/api             # Application entry point
├── internal/           # Private business logic
│   ├── database        # Database connection and configuration
│   ├── handlers        # HTTP controllers (Request handling)
│   ├── models          # Data structure definitions
│   └── repository      # Persistence layer (SQL)
├── scripts/            # Useful scripts (DB, setups)
├── tests/              # Automated tests
└── docs/               # Documentation and assets
```

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/health` | Check API health status |
| `GET` | `/api/tasks` | List all tasks |
| `POST` | `/api/tasks` | Create a new task |
| `GET` | `/api/task?id={id}` | Get a task by ID |
| `PUT` | `/api/task?id={id}` | Update an existing task |
| `DELETE` | `/api/task?id={id}` | Delete a task |

## Quick Start

1. **Requirements**: Have Go 1.21+ and PostgreSQL installed.
2. **Configuration**:

   ```bash
   cp .env.example .env
   # Edit the .env file with your database credentials
   ```

3. **Run**:

   ```bash
   make run
   ```

4. **Test**:

   ```bash
   make test
   ```
