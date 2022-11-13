CREATE TABLE note_has_tag
(
    note_id BINARY(16) NOT NULL,
    tag_id  BINARY(16) NOT NULL,
    PRIMARY KEY (note_id, tag_id),
    CONSTRAINT note_has_tag_note_id_fk FOREIGN KEY (note_id) REFERENCES note (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT note_has_tag_tag_id_fk FOREIGN KEY (tag_id) REFERENCES note_tag (id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;
