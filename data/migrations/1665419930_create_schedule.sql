CREATE TABLE schedule
(
    id           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    timetable_id INT UNSIGNED NOT NULL,
    is_even      TINYINT(1)   NOT NULL,
    title        VARCHAR(127) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT timetable_id_fk FOREIGN KEY (timetable_id) REFERENCES timetable (id)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
