CREATE TABLE migration_versions
(
    version BIGINT(20) NOT NULL,
    date    DATE       NOT NULL DEFAULT NOW(),
    PRIMARY KEY (version)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
