# Go Not Another Task Manager (Go-NATM)

A simple task management app that allows you users to log time against issues.  Golang back-end with React front-end.

## Install

1. Have a connection string for a database handy
  - Only tested with Postgres, but others may work fine.
2. Copy `.env.example` to a new `.env` setting the missing credentials.
  - Google OAuth is mandatory for now, but will likely have configurable auth strategies later.
3. Use Golang-Migrate command to initialize the DB
  - `migrate -source file://migrations -database postgres://username:pw@localhost:5432/gonatm up`