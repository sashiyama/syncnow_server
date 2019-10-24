CREATE TABLE IF NOT EXISTS user_bonds (
    id uuid NOT NULL REFERENCES users(id),
    relation_id uuid NOT NULL REFERENCES users(id),
    created_at timestamp without time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp without time zone NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY(id, relation_id)
);

CREATE TRIGGER update_tri BEFORE UPDATE ON user_bonds FOR EACH row EXECUTE PROCEDURE set_update_time();
