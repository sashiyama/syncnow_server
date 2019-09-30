CREATE TABLE IF NOT EXISTS user_refresh_tokens (
    id uuid DEFAULT public.gen_random_uuid() PRIMARY KEY NOT NULL,
    user_access_token_id uuid NOT NULL REFERENCES user_access_tokens(id),
    token uuid DEFAULT public.gen_random_uuid() NOT NULL UNIQUE,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp without time zone NOT NULL DEFAULT current_timestamp
);

