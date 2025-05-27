CREATE SCHEMA IF NOT EXISTS participant;

CREATE TABLE IF NOT EXISTS participant.participants (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    schedule_ids JSONB DEFAULT '[]',
    status TEXT NOT NULL
)