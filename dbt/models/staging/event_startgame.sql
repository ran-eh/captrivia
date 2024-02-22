{{ config(materialized='view') }}

select 
    timestamp,
    event_id,
    session_id,
    program
from {{ source('captrivia', 'events') }}
where type = 'start_game'