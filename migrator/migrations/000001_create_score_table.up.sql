CREATE TABLE IF NOT EXISTS score(
    score_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    score_value INT NOT NULL
);