## Project Overview

[ТЗ](docs/ТЗ.md)


This project is a Go-based web service named `denet-test`. It appears to be a user management service, with features related to users, tasks, and leaderboards. The service is containerized using Docker and uses a PostgreSQL database for data storage. The architecture follows a clean structure, separating concerns into different packages like `repo`, `usecase`, and `api`.

**Key Technologies:**

*   **Language:** Go
*   **Web Framework:** Gin
*   **Database:** PostgreSQL
*   **Database Migration:** golang-migrate
*   **Containerization:** Docker, Docker Compose

## Building and Running

The primary way to build and run the project is by using the provided `Makefile` and Docker Compose.

**Prerequisites:**

*   Docker
*   Docker Compose


**Commands:**

*   **Start the application:**
    ```bash
    make up
    ```
    This command will build the Docker images and start the `denet-users-service` and `postgres` containers in the background.

*   **Stop the application:**
    ```bash
    make down
    ```
    This command will stop and remove the running containers.


# Development Conventions

## OpenAPI documentation

[swagger.yaml](docs/swagger.yaml)
*   `http://localhost:8080/api/v1/swagger/index.html`

## API Endpoints
The following API endpoints are available:

*   GET /users/{id}/status - вся доступная информация о пользователе


*   GET /users/leaderboard - топ пользователей с самым большим балансом

*   POST /users/{id}/task/complete - выполнение задания

*   POST /users/{id}/referrer - ввод реферального кода (может быть id другого пользователя)

## Database Migrations

Database migrations are located in the `./migrate/migrations` directory. To run the migrations, the application will automatically run them on startup.