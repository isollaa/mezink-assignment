CREATE TABLE IF NOT EXISTS records (
    id smallserial NOT NULL,
    name varchar(255) NOT NULL,
    marks integer[] NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
PRIMARY KEY (id)
);

CREATE INDEX idx_records_name ON records (name);