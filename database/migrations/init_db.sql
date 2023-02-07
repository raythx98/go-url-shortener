CREATE DATABASE go_url_shortener; -- done by DBA

use go_url_shortener;

CREATE TABLE links (
    link_id        BIGINT      AUTO_INCREMENT PRIMARY KEY ,
    link           TEXT        NOT NULL,
    shortened_link VARCHAR(12) NOT NULL
);

CREATE INDEX shortened_link_idx
    ON links (shortened_link);

CREATE USER 'goserver'@'localhost' -- done by DBA
    IDENTIFIED WITH mysql_native_password BY 'password';
GRANT ALL ON go_url_shortener.* -- done by DBA
    TO 'goserver'@'localhost';