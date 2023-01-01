CREATE TABLE image
(
    id   BINARY(16)   NOT NULL,
    path VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
