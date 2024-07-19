-- name: GetLatestMetrics :one
SELECT 
    DATE(timestamp) as date,
    SUM(keypresses) as total_keypresses,
    SUM(mouse_clicks) as total_mouse_clicks,
    CAST(SUM(mouse_distance_in) AS DOUBLE PRECISION) as total_mouse_travel_in,
    CAST(SUM(mouse_distance_mi) AS DOUBLE PRECISION) as total_mouse_travel_mi,
    SUM(scroll_steps) as total_scroll_steps
FROM metrics
WHERE DATE(timestamp) = (
    SELECT DATE(timestamp)
    FROM metrics
    ORDER BY timestamp DESC
    LIMIT 1
)
GROUP BY DATE(timestamp);
