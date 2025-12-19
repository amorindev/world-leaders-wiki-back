# World Leaders Wiki – Backend

Backend API for World Leaders Wiki, a platform that centralizes structured information about political, social, and historical leaders around the world.

This project is designed with scalability and clean architecture in mind, allowing easy integration with web frontend.

Frontend

The web frontend for this project is available here:
- https://github.com/amorindev/world-leaders-wiki-front


## Architecture

- Language: Go (Golang)
- Architecture: Arq Hexagonal
- Database: MongoDB
- Storage: Object storage (MinIO / S3 compatible)
- API Style: REST
- Environment-based configuration

## Installation

1. Clone the repository:
    ```bash
    git clone github.com/amorindev/world-leaders-wiki-back
    cd world-leaders-wiki-back
    ```

2. Download dependencies:
    ```bash
    go mod tidy
    ```

3. Set environment variables, add a `.env` file based on `env.example`:
4. Run the project:
    ```bash
    make run 
    ```
   Or:
    ```bash
    go run main.go
    ```
