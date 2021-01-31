#!/bin/bash
goose -dir ./migrations mysql "db:db@/db?parseTime=true" up
# usage: https://github.com/mattes/migrate