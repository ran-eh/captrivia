CREATE TABLE IF NOT EXISTS public.events
(
    "timestamp" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    program character varying COLLATE pg_catalog."default",
    type character varying COLLATE pg_catalog."default",
    data jsonb,
    context jsonb
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.events
    OWNER to postgres;