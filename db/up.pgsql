CREATE TABLE IF NOT EXISTS public.events
(
    "timestamp" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    event_id uuid DEFAULT gen_random_uuid(),
    session_id character varying COLLATE pg_catalog."default",
    program character varying COLLATE pg_catalog."default",
    type character varying COLLATE pg_catalog."default",
    data jsonb,
    context jsonb
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.events
    OWNER to postgres;