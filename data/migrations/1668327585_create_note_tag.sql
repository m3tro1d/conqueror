CREATE TABLE note_tag
(
    id      BINARY(16)   NOT NULL,
    name    VARCHAR(200) NOT NULL,
    user_id BINARY(16)   NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT note_tag_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
