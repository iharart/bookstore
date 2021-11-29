CREATE DATABASE IF NOT EXISTS bookstore;
GRANT ALL PRIVILEGES ON bookstore.* TO 'root'@'%' IDENTIFIED BY 'mysql';
USE bookstore;
CREATE TABLE Genre
(
    id                          bigint(20)   UNSIGNED NOT NULL,
    name                        varchar(99)  NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE Book
(
    id                          bigint(20)   UNSIGNED NOT NULL,
    name                        varchar(99)  NOT NULL,
    genre_id                    bigint(20)   UNSIGNED NOT NULL,
    price                       double(10,2) UNSIGNED NOT NULL,
    amount                      bigint(20)   UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (genre_id)  REFERENCES Genre (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO Genre
VALUES (1, 'Adventure');
INSERT INTO Genre
VALUES (2, 'Classics');
INSERT INTO Genre
VALUES (3, 'Fantasy');

INSERT INTO Book
VALUES (1, 'Don Quixote', 1, 100, 5);
INSERT INTO Book
VALUES (2, 'Moby Dick', 2, 200, 10);
INSERT INTO Book
VALUES (3, 'Game of thrones', 3, 300, 10);
INSERT INTO Book
VALUES (4, 'Dracula', 2, 300, 20);
