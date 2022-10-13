CREATE TABLE lesson
(
    id                 BINARY(16)  NOT NULL,
    lesson_interval_id BINARY(16)  NOT NULL,
    subject_id         BINARY(16)  NOT NULL,
    type               TINYINT(1)  NOT NULL,
    auditorium         VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT lesson_lesson_interval_id_fk FOREIGN KEY (lesson_interval_id) REFERENCES lesson_interval (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT lesson_subject_id_fk FOREIGN KEY (subject_id) REFERENCES subject (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
