ALTER TABLE user
    ADD COLUMN avatar_id BINARY(16) DEFAULT NULL AFTER password,
    ADD CONSTRAINT user_avatar_id_fk FOREIGN KEY (avatar_id) REFERENCES image (id)
        ON UPDATE CASCADE ON DELETE SET NULL;
