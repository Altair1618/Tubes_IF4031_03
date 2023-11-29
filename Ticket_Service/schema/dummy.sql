INSERT INTO events (id, event_name, event_time, event_location, created_at, updated_at)
VALUES (
    '00000000-0000-0000-0000-000000000001',
    'Konser Miku',
    NOW() + INTERVAL '30 days',
    'Jakarta',
    NOW(),
    NOW()
);

INSERT INTO tickets (id, price, event_id, seat_id, status, created_at, updated_at)
SELECT
    '00000000-0000-0000-0000-000000000001',
    50,
    (SELECT id FROM events WHERE event_name = 'Konser Miku'),
    chr(65 + (i / 1000) % 26) || chr(65 + (i / 100) % 26) || 
    CONCAT(
        CASE WHEN i % 1000 < 100 THEN '0' END,
        CASE WHEN i % 1000 < 10 THEN '0' END,
        i % 1000
    ),
    'ON GOING',
    NOW(),
    NOW()

FROM generate_series(1, 1) i;

INSERT INTO tickets (id, price, event_id, seat_id, status, created_at, updated_at)
SELECT
    gen_random_uuid(),
    50,
    (SELECT id FROM events WHERE event_name = 'Konser Miku'),
    chr(65 + (i / 1000) % 26) || chr(65 + (i / 100) % 26) || 
    CONCAT(
        CASE WHEN i % 1000 < 100 THEN '0' END,
        CASE WHEN i % 1000 < 10 THEN '0' END,
        i % 1000
    ),
    'OPEN',
    NOW(),
    NOW()
FROM generate_series(2, 30) i;