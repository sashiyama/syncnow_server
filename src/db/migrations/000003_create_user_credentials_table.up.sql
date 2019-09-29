CREATE TABLE IF NOT EXISTS user_credentials (
    id uuid DEFAULT public.gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id),
    email character varying NOT NULL UNIQUE,
    password_digest character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
