BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create function trigger for automatically set updated_at
CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

COMMIT;