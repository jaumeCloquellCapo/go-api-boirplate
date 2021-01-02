# Go API Boilerplate

When I want to start to build Go API project, i don't have a good solid base to start and usually I add the library and add another required thing one by one along the time, and then change again if I find another better library or another better way to do thing. So I tried to research architecture, library and software component/layer that I think better suits to be included for solid golang project.

## Objectives
* [x] Scalable, must be able to run more than one instance.
* Dockerized, runnable on minikube.
* Unit tested, must be able to run "go test ./..." directly from clone.
* [x] Integration tested, recommend docker-compose.
* OpenAPI/Swagger (or similar for gRPC) documented.
* [x] dep vendored, but using the standard library often, instead of piling on dependencies.
* [x] Authenticated and Authorized via apikeys and/or user accounts.
* Built and tested via CI: Travis, CircleCi, etc. Recommend Makefile for task documentation.
* [x] Flag & ENV config, API keys, ports, dev mode, etc.
* "why" comments, not "what" or "how" which should be clear through func/variable names and godoc comments.
* Use of Context to limit request time.
* Leveled JSON logging, logrus or similar.
* [x] Postgres/MySQL, sqlx or an ORM.
* [x] Redis/memcache for scalable caching.
* Well documented README.md with separate sections for API user and service developer audiences. Maybe even include graphviz or mermaidJS UML diagrams.
* Clean git history with structured commits and useful messages. No merge master commits.
* Passing go fmt, go lint, or better, go-metalinter in the CI.

## Commands

go run main.go server

### Authentication
For passwordless login following routes are available:

Path | Method | Required JSON | Header | Description
---|---|---|---|---
/login | POST |  emial, password | | the email you want to login with (see below)
/auth/users | GET |  | Authorization: "Bearer refresh_token"  | 
/auth/users/$id | GET | | Authorization: "Bearer refresh_token"  | 
/auth/register | POST | | Authorization: "Bearer refresh_token" | Create new user
/auth/logout | POST | | Authorizaiton: "Bearer refresh_token" | logout from this device
