# Dating Apps
### Structure Layer
```
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md [you are here!]
├── wire.go
├── wire_gen.go
├── bin
├── configs
├── docs
├── http
├── infras
├── migrations
├── shared
└── internal
          ├── domains
          │   ├── dating
          │   │   ├── model
          │   │   │   ├── dto
          │   │   ├── repository
          │   │   └── service
          │   └── user
          │       ├── model
          │       │   ├── dto
          │       ├── repository
          │       └── service
          └── handlers
              ├── dating
              └── user

```
- docker-compose.yml : service configuration for postgres and redis dev with docker
- go.mod & go.sum : dependency management in Go projects
- main.go : entry point for a Go application
- Makefile : utility to automate the building and management of projects. It contains a set of rules and commands that specify how to run, test, migration, linter etc
- wire.go & wire_gen.go : ependency injection (DI) tool for Go that helps generate code for wiring together dependencies in a type-safe manner
- bin/ :  store compiled executable files or binaries
- configs/ : save and manage configuration in Go projects
- docs/ : documentation how to use API endpoint (swagger)
- http/ : manage http router and middleware
- infras/ : init connection for redis and postgresql
- migrations/ : manage migration for database schema
- shared/ : manage helper, util or third party support
- internal/ : logic for business flow

#### Business Flow Structure Layer
- domains/ : split entity for avoid bloated data
- handlers/ : manage router and validation
- service/ :  
  - component or module in a Go project that encapsulates a specific set of functionalities or operations
  - represents a higher-level abstraction that handles a specific domain or business logic
- repository/ :
  - provides methods or functions to perform Data Access and Data Mapping
  - interact directly with infras (redis & postgres)
- model/ :
  - defining Data Structure from Repository
- dto/ :
  - defining Validation and Data Mapping
  - defining response data or request data body
  - methods or functions for converting their data to external representations, such as JSON, XML, or protocol buffers, and vice versa

## Setup and Installation

To set up and run this Go project locally, please follow these steps:
### Prerequisites
- Go Programming Language
- Docker

### Run Project

1. Turn on the development infra with docker
```shell
make docker_dev
```
2. Run migration for schema database (you can change MIGRATION_STEP on Makefile for amount of step run in once command)
```shell
make migrate_up
```
3. Run development server
```shell
make dev
```
4. open localhost with port you have configured at .env and swagger path to see the api documentation, example: http://localhost:8080/swagger/index.html
5. Viola!