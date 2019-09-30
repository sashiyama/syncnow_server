CREATE TABLE IF NOT EXISTS user_profiles (
    id uuid DEFAULT public.gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id),
    display_name character varying NOT NULL UNIQUE,
    name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp without time zone NOT NULL DEFAULT current_timestamp
);
