INSERT INTO score(user_id, created_at, score_value)
VALUES
    (1, NOW()::timestamp, 3000),
    (2, NOW()::timestamp, 2000),
    (3, NOW()::timestamp, 1000);