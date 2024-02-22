{{ config(materialized='view') }}

select 
    timestamp,
    event_id,
    session_id,
    program,
    cast(data->'Answer'->>'questionId' as numeric) as questionId,
    cast(data->'Answer'->>'answer' as numeric) as answer,
    cast(data->>'Correct' as boolean) as Correct
from {{ source('captrivia', 'events') }}
where type = 'answer'