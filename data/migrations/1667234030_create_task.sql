CREATE TABLE task
(
    id          BINARY(16)    NOT NULL,
    user_id     BINARY(16)    NOT NULL,
    due_date    DATE          NOT NULL,
    title       VARCHAR(200)  NOT NULL,
    description VARCHAR(1000) NOT NULL,
    status      TINYINT(1)    NOT NULL,
    subject_id  BINARY(16),
    PRIMARY KEY (id),
    CONSTRAINT task_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT task_subject_id_fk FOREIGN KEY (subject_id) REFERENCES subject (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
