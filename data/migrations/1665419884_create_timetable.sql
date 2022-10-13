CREATE TABLE timetable
(
    id      BINARY(16) NOT NULL,
    user_id BINARY(16) NOT NULL,
    type    TINYINT(1) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT timetable_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
