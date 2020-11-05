# Go API Boilerplate

When I want to start to build Go API project, i don't have a good solid base to start and usually I add the library and add another required thing one by one along the time, and then change again if I find another better library or another better way to do thing. So I tried to research architecture, library and software component/layer that I think better suits to be included for solid golang project.

## Architecture
This project follows SOLID & Clean architecture

## API Routes

### Authentication
For passwordless login following routes are available:

Path | Method | Required JSON | Header | Description
---|---|---|---|---
/auth/login | POST | email | | the email you want to login with (see below)
/auth/token | POST | token | | the token you received via email (or printed to stdout if smtp not set)
/auth/refresh | POST | | Authorization: "Bearer refresh_token" | refresh JWTs
/auth/logout | POST | | Authorizaiton: "Bearer refresh_token" | logout from this device
