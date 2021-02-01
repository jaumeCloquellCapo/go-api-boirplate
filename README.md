# Go API CRUD

Go API CRUD app architecture showcase using Gin, Mysql and redis. 

# Architecture Overview #
The app is designed to use a layered architecture. The architecture is heavily influenced by the Clean Architecture and Hexagonal Architecture. [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) is an architecture where `the business rules can be tested without the UI, database, web server, or any external element`.

<p align="center">
  <img src="https://cdn-images-1.medium.com/max/719/1*ZNT5apOxDzGrTKUJQAIcvg.png" width="350"/>
</p>

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
├── internal                            # Contains all application packages
      ├── helpers                      
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
├── tests                               # End-To-End test  
├── mocks                               # Fake Service to test controllers handler functions
```

## Implementation Notes

### Storage providers

I picked Mysql and Redis because on the one hand, MySQL offers good performance and it is used to save relational data how for example the users table. On the other hand, Redis is used as an external cache to save the session info of the users. Although is only used as a session cache, it could be a data store to cache request o any type of information.

In this example we will connect to a Mysql database, but the syntax (minus some small SQL semantics) is the same for a SQL or PostgreSQL database.


    // DbStore ...
    type DbStore struct {
    *sql.DB
    }

    // Opening a storage and save the reference to `Database` struct.
    func InitializeDB() *DbStore {
        //dataSourceName := fmt.Sprintf(core.Database.Username + ":" + core.Database.Password + "@/" + core.Database.Database)
        cnf := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
    
        var err error
        var db *sql.DB
    
        if db, err = sql.Open("mysql", cnf); err != nil {
            log.Fatal(err)
        }
    
        retryCount := 30
        for {
            err := db.Ping()
            if err != nil {
                if retryCount == 0 {
                    log.Fatalf("Not able to establish connection to database")
                }
    
                log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retryCount))
                retryCount--
                time.Sleep(2 * time.Second)
            } else {
                break
            }
        }
    
    
        if errPing := db.Ping(); errPing != nil {
            log.Fatal(errPing)
        }
        return &DbStore{
            db,
        }
    }

I added a simple for functions to retry the connection for 30 times.

### Dependency Injection

As a software developer split our code into different layers is a requirement if we desire to make it clean and maintainable.
Usually, the boundaries are placed at least between infrastructure and business logic. When we are dealing specially with complex business logic, it is desirable that infrastructure depends on our business logic, so that we don’t break our software when changing the infrastructure.
The first decision when developing a new software project is to materialize this layer split by choosing an architecture. Most of the time I choose Clean Architecture, but you have another good option like Domain-Driven Design.
Independently of the architecture you choose, we have to glue the pieces from the different layers to come up with a new feature and this is where Dependency Injection shines.

### Tests

Testing is an important part of any application. There are two approaches we can take to testing Go web applications. The first approach is a unit-test style approach. The other is more of an end-to-end approach. In this chapter we'll cover both approaches.

### Unit Tests

When testing a component, we ideally want to isolate it completely to avoid having failures elsewhere to compromise our tests. This is especially harder when the component we want to test has dependencies on other components from different layers in our software. In the scenario we are using here, our service implementation depends on a component from the UserService or AuthService layers to access information about the users.
To promote the desired isolation, it is common for developers to write fake simplified implementations of those dependencies to be used during the tests. Those fake implementations are called mocks.
We can create a mock implementation of the UserService to be injected into the controller implementation for the test execution.

    func (m *MockUserService) FindByEmail(email string) (user *model.User, err error) {
        return &model.User{
            ID:         0,
            Name:       "0",
            LastName:   "0",
            Password:   nil,
            Email:      "0",
            Country:    "0",
            Phone:      "0",
            PostalCode: "0",
        }, nil
    }

To enable the test execution, it is necessary that the mock provides a behavior compatible with all test cases we want to validate, otherwise we cannot achieve the desired test coverage.
### Unit End-to-end
End to end allows us to test applications through the whole request cycle. Where unit testing is meant to just test a particular function, end to end tests will run the middleware, router, and other that a request my pass through.

The folder tests have all e2e tests 


## Commands

### Development on local

#### Manually fetch dependencies from go.mod?
    go mod download
#### Run server on local
    make dev_up && make migrate_dev
#### Run server on local
    make run

### Dockerized api

#### Generate docker image
    make pro_up

If you received this follow error when docker-compose up is running, you only have to rerun the same command ( make pro_up )

    FileNotFoundError: [Errno 2] No such file or directory: '/tmp/tmpxd1tbhth'
    [114003] Failed to execute script docker-compose
    make: *** [Makefile:21: pro_up] Error 255

#### Run migrate into docker shell

After of this, is necessary access to docker shell to run 
    
    make migrate_pro


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
