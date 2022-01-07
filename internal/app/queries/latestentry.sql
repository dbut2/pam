SELECT e.id, eh.date, eh.entry FROM entry AS e

LEFT JOIN entry_history eh on eh.id = e.id
LEFT JOIN metadata m on m.id = eh.metadata

WHERE m.timestamp = (
    SELECT MAX(m.timestamp)
    FROM entry_history AS latest
    LEFT JOIN metadata AS m ON latest.metadata = m.id
    WHERE eh.id = latest.id
)

AND m.author = ?

ORDER BY eh.date DESC, m.timestamp DESC