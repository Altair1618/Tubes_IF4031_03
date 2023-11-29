INSERT INTO booking_history (ticket_id, status, user_id, created_at)
SELECT
    '00000000-0000-0000-0000-000000000001' AS ticket_id,
    'IN QUEUE'::booking_status AS status,
    '6rqxaaoovfmn558'::varchar(15) AS user_id,
    NOW() - (i * INTERVAL '1 hour') AS created_at
FROM generate_series(1, 4) i;

INSERT INTO booking_history (ticket_id, status, user_id, created_at)
SELECT
    '00000000-0000-0000-0000-000000000001' AS ticket_id,
    'WAITING FOR PAYMENT'::booking_status AS status,
    '6rqxaaoovfmn558'::varchar(15) AS user_id,
    NOW() - (i * INTERVAL '1 hour') AS created_at
FROM generate_series(5, 5) i;