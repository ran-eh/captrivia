select
	session_id,
	s.timestamp as started_at,
	e.timestamp as ended_at,
	e.final_score as score,
	count(*) as num_answered,
	count(*) filter (where a.correct) as num_correct
from
	{{ ref('event_startgame') }} as s
	left join {{ ref('event_endgame') }} as e using(session_id)
	left join {{ ref('event_answer') }} as a using(session_id)
group by 1, 2, 3, 4
