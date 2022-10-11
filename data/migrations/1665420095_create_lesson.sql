CREATE TABLE lesson
(
    id                 BINARY(16)   NOT NULL,
    lesson_interval_id INT UNSIGNED NOT NULL,
    subject_id         INT UNSIGNED NOT NULL,
    type               TINYINT(1)   NOT NULL,
    auditorium         VARCHAR(50)  NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT lesson_interval_id_fk FOREIGN KEY (lesson_interval_id) REFERENCES lesson_interval (id),
    CONSTRAINT subject_id_fk FOREIGN KEY (subject_id) REFERENCES subject (id)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
