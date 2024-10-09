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
