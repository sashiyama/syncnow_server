CREATE TABLE IF NOT EXISTS user_coordinates (
    id uuid DEFAULT public.gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id),
    coordinates public.geometry(Point,4326) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
CREATE INDEX index_user_coordinates_on_coordinates ON user_coordinates USING gist (coordinates);
