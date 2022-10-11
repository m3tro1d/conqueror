CREATE TABLE user
(
    id       BINARY(16)  NOT NULL,
    login    VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX login_idx (login)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
