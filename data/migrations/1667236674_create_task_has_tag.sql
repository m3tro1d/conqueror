CREATE TABLE task_has_tag
(
    task_id BINARY(16) NOT NULL,
    tag_id  BINARY(16) NOT NULL,
    PRIMARY KEY (task_id, tag_id),
    CONSTRAINT task_has_tag_task_id_fk FOREIGN KEY (task_id) REFERENCES task (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT task_has_tag_tag_id_fk FOREIGN KEY (tag_id) REFERENCES task_tag (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
