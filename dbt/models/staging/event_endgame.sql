{{ config(materialized='view') }}

select 
    timestamp,
    event_id,
    session_id,
    program,
    cast(data->>'finalScore' as numeric) as final_score
from {{ source('captrivia', 'events') }}
where type = 'end_game'