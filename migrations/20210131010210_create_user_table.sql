-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id int(11)  UNSIGNED NOT NULL AUTO_INCREMENT,
    last_name varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    country varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    postal_code varchar(255) NOT NULL,
    phone varchar(255) NOT NULL,
    PRIMARY KEY (id) ,
    KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
