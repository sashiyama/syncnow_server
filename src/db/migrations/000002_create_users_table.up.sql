CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT public.gen_random_uuid() PRIMARY KEY NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
