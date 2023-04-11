DROP DATABASE IF EXISTS test;
CREATE DATABASE test;
\c test;

CREATE TABLE test (
    a varchar(255) NOT NULL,
    b varchar(255) NOT NULL,
    c varchar(255) NOT NULL
);
