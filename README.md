# Book The Flight
## Prerequisite

- Docker
- Docker compose

## Quick start

**Run servers using docker-compose**
- Backend (Go)
```
cd backend
docker-compose down -v && docker-compose up --build
```
- Frontend (React JS)
```
cd frontend/booktheflight
docker-compose down -v && docker-compose up --build
```