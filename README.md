# Go API Boilerplate

Easily extendible RESTful API boilerplate aiming to follow idiomatic go and best practice.

The goal of this boiler is to have a solid and structured foundation to build upon on.

Any feedback and pull requests are welcome and highly appreciated. Feel free to open issues just for comments and discussions.

## Features
The following feature set is a minimal selection of typical Web API requirements:

- Configuration using [viper](https://github.com/spf13/viper)
- CLI features using [cobra](https://github.com/spf13/cobra)
- PostgreSQL support including migrations using [go-pg](https://github.com/go-pg/pg)
- Structured logging with [Logrus](https://github.com/sirupsen/logrus)
- Routing with [chi router](https://github.com/go-chi/chi) and middleware
- JWT Authentication using [jwt-go](https://github.com/dgrijalva/jwt-go) with example passwordless email authentication
- Request data validation using [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- HTML emails with [gomail](https://github.com/go-gomail/gomail)

## Start Application
- Clone this repository
- Create a postgres database and set environment variables for your database accordingly if not using same as default
- First initialize the database running all migrations found in ./database/migrate at once with command *migrate*: ```go run main.go migrate```
- Run the application with command *serve*: ```go run main.go serve```

## API Routes

### Authentication
For passwordless login following routes are available:

Path | Method | Required JSON | Header | Description
---|---|---|---|---
/auth/login | POST | email | | the email you want to login with (see below)
/auth/token | POST | token | | the token you received via email (or printed to stdout if smtp not set)
/auth/refresh | POST | | Authorization: "Bearer refresh_token" | refresh JWTs
/auth/logout | POST | | Authorizaiton: "Bearer refresh_token" | logout from this device
