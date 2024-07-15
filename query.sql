-- name: GetLatestMetrics :one
SELECT 
    DATE(timestamp) as date,
    SUM(keypresses) as total_keypresses,
    SUM(mouse_clicks) as total_mouse_clicks
FROM metrics
WHERE DATE(timestamp) = (
    SELECT DATE(timestamp)
    FROM metrics
    ORDER BY timestamp DESC
    LIMIT 1
)
GROUP BY DATE(timestamp);
