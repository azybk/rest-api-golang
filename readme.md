CREATE DATABASE PERPUSTAKAAN;

CREATE TABLE CUSTOMERS
(
id varchar(100),
code varchar(10),
name varchar(100),
created_at timestamp DEFAULT NULL,
updated_at timestamp DEFAULT NULL,
deleted_at timestamp DEFAULT NULL,
PRIMARY KEY (id)
);

INSERT INTO customers VALUES(gen_random_uuid(), 'A-001', 'aink tea', '2020-10-09 23:26:07', '2020-10-09 23:26:10', '2020-10-09 23:26:12');

go get -u github.com/golang-jwt/jwt/v5 github.com/gofiber/contrib/jwt

CREATE TABLE public.users (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);
