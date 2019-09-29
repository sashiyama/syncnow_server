CREATE TABLE IF NOT EXISTS user_access_tokens (
    id uuid DEFAULT public.gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id),
    token uuid DEFAULT public.gen_random_uuid() NOT NULL UNIQUE,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
