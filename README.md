# mezink-assignment

### Instructions

1. Clone repo
2. Copy and paste `.env.example` to `.env` and adjust based on your configuration
3. Run service : 
- local
```
go run .
```
- docker
```
make docker_start
```

### Database Connection

This service requiring pre-defined **_PostgreSQL Database_** to connect and perform the database migration.

**Note:** *make sure to set these variables properly*
```
INFRA_POSGRE_DB_NAME
INFRA_POSGRE_HOST
INFRA_POSGRE_PASSWORD
INFRA_POSGRE_PORT
INFRA_POSGRE_TIMEZONE
INFRA_POSGRE_USERNAME
```

Set `INFRA_POSGRE_ENABLE_MIGRATION` as `true` to enable migration feature.

**Additional information:** *this service is using [go-migrate](https://github.com/golang-migrate/migrate) for the migration tool.*

## Technical Information

For technical information about this service, please visit the following links:

* [Product Requirement](https://docs.google.com/document/d/1vXWU5loXw0yukNEC3g0mhyJghCBl5Y6t1CJWqjEtvZc/edit?usp=sharing)
* [Postman Documenter](https://documenter.getpostman.com/view/10609164/2sA3BhduVq)
