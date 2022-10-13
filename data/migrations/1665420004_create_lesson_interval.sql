CREATE TABLE lesson_interval
(
    id          BINARY(16) NOT NULL,
    schedule_id BINARY(16) NOT NULL,
    weekday     TINYINT(1) NOT NULL,
    start_time  TIME       NOT NULL,
    end_time    TIME       NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT lesson_interval_schedule_id_fk FOREIGN KEY (schedule_id) REFERENCES schedule (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
