CREATE TABLE IF NOT EXISTS url_statuses (
    id SERIAL PRIMARY KEY,
    url_id INTEGER REFERENCES monitored_urls(id) ON DELETE CASCADE,
    status_code INTEGER,
    is_up BOOLEAN,
    checked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
