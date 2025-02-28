CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       task_name VARCHAR(255) NOT NULL,
                       is_done BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);