CREATE TABLE schedule
(
    id           BINARY(16)   NOT NULL,
    timetable_id BINARY(16)   NOT NULL,
    is_even      TINYINT(1)   NOT NULL,
    title        VARCHAR(127) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT schedule_timetable_id_fk FOREIGN KEY (timetable_id) REFERENCES timetable (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
