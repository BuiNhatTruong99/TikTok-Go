
## Project Structure
    tiktok-go
    │
    ├── .github                 #This directory contains the github action for run automated tests
    │   └── workflows
    ├── cmd                     #This directory contains the entry point for application.
    │   └── root.go
    │
    ├── composer                #This directory for compose and set up api sevices 
    │   └── service_composer.go
    |
    ├── config                  #Configuration files and code related to configuration
    │   └── config.go
    │
    ├── midlleware              #This directory contains midlldewares
    |
    ├── internal                #This is where most of application code resides.
    │   └── auth.go
    │   └── post.go
    │   └── session.go
    │
    ├── migrations              #Database migration files.
    │   └── ...
    │
    ├── pkg                     #External packages and utilities that can be shared across different projects
    │   └── ...
    │
    └── main.go                 #The entry point of application


## Setup local development

### Install tools
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Docker desktop](https://www.docker.com/products/docker-desktop)
  + pull postgres image
       ```bash
      docker pull postgres:16-aline
      ```
### Setup infrastructure
- Start postgres container:
    ```bash
    make postgres
    ```
- Create tiktok database:
    ```bash
    make createdb
    ```
- Run db migrate-user:
    ```bash
    make migrate-user
    ```
- Run db migration up:
    ```bash
    make migrateup

### How to run
- Run mockgen:
    ```bash
    make mock
    ```
- Run test:
    ```bash
    make test
    ```
- Run app:
    ```bash
    make run
    ```