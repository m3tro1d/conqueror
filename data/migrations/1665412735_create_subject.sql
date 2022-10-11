CREATE TABLE subject
(
    id      BINARY(16)   NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    title   VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES user (id)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
