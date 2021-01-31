# Go API CRUD

Go API CRUD app architecture showcase using Gin, PostgreSQL and redis. 

# Architecture Overview #
The app is designed to use a layered architecture. The architecture is heavily influenced by the Clean Architecture and Hexagonal Architecture. [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) is an architecture where `the business rules can be tested without the UI, database, web server, or any external element`.

<p align="center">
  <img src="https://cdn-images-1.medium.com/max/719/1*ZNT5apOxDzGrTKUJQAIcvg.png" width="350"/>
</p>

## Implementation Notes

### Storage providers
I picked Mysql and Redis because for one hand, mysql offer a good performance and it is used to save relational data how for examples the users table. On the other hand, redis is used as a external cache to save session info of the users. Although is only used like a session cache, it could will be a data store to cache request o any type of information.

### Dependency Injection
As a software developer split our code in different layers is a requirement if we desire to make it clean and maintainable.
Usually the boundaries are placed at least between infrastructure and business logic. When we are dealing specially with complex business logic, it is desirable that infrastructure depends on our business logic, so that we don’t break our software when changing the infrastructure.
The first decision when developing a new software project, is to materialize this layer split choosing an architecture. Most of the times I choose Clean Architecture, but you have another good options like Domain Driven Design.
Independently from the architecture you choose, we have to glue the pieces from the different layers to come up with a new feature and this is where Dependency Injection shines.


In our project the dependencies are provided through a class constructor. This states that a class should not configure its dependencies statically but should be configured by some other class from outside.  It is the fifth principle of S.O.L.I.D — the five basic principles of object-oriented programming and design by Uncle Bob — which states that a class should depend on abstraction and not upon concretions (in simple terms, hard-coded).


## Objectives
* [x] Scalable, must be able to run more than one instance.
* [x] Dockerized
* [x] Unit tested, must be able to run "go test ./..." directly from clone.
* [x] Integration tested, recommend docker-compose.
* [x] dep vendored, but using the standard library often, instead of piling on dependencies.
* [x] Authenticated via apikeys and/or user accounts.
* Built and tested via CI: Travis, CircleCi, etc. Recommend Makefile for task documentation.
* [x] Flag & ENV config, API keys, ports, dev mode, etc.
* [x] Use of Context to limit request time.
* [x] Leveled JSON logging, logrus or similar.
* [x] MYSQL
* [x] Redis/memcache for scalable caching.
* [x] Passing go fmt, go lint, or better, go-metalinter in the CI.

## Project structure

```sh
stygis/
├── bin     
├── app                                 # Domain packages are here, contains business logic and interfaces that belong to each domain
      ├── controller                    # handler for rest API technology
      ├── model 
            ├── users                   # only sample domain, user package which handler for user business logic  
      ├── repository                    # this is the only file to declare interface methods from storage and repository. also where to put func init the package.
├── cmd                                 # Packages that provide support for a specific program that is being built, this contains cmd for REST
├── internal                            # Contains all application packages
      ├── helpers                       # For declaring all constants variable based on it's entity, declaring with full name descripting it
            ├── encryption              # you can have any encryption method here
      ├── route                         # where the routing for handlers are assigned based on method and url
      ├── dic                           # dependency injection container
      ├── middleware                    
             ├── auth                   # middleware control of the access of unauthenticated users
             ├── cors                   # middleware that can be used to enable CORS with various options
      ├── storage                       # this is where you put the data storing code. whether persistence like postgresql etc. and caching like redis, etc. 
            ├── cache                   # contains functions to open database redis connection
            ├── persistence             # contains functions to open database mysql connections
├── migrations                          # Contains sql files to migrate database
```

## Commands

### Run server
make run

### Build Server
make build

### Migrate database
make build

## Routes

Path | Method  | Header | Description
---|---|---|---
/login | POST |   the email you want to login with (see below)
/signup | POST |    the email you want to login with (see below)
/auth/logout | POST   | Authorizaiton: "Bearer refresh_token" | logout from this device
/auth/users | GET  | Authorization: "Bearer refresh_token"  | 
/auth/users/:id | GET | Authorization: "Bearer refresh_token"  | Get details
/auth/users/:id | PUT  | Authorizaiton: "Bearer refresh_token" | Update
/auth/users/:id | DELETE  | Authorizaiton: "Bearer refresh_token" | Remove
