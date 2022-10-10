CREATE TABLE user
(
    id       INT UNSIGNED NOT NULL AUTO_INCREMENT,
    login    VARCHAR(50)  NOT NULL,
    password VARCHAR(50)  NOT NULL,
    nickname VARCHAR(50)  NOT NULL,
    PRIMARY KEY (id),
    INDEX login_idx (login)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
