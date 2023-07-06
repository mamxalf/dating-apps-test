BEGIN;

CREATE TABLE IF NOT EXISTS swipes (
                         id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
                         user_id UUID NOT NULL,
                         profile_id UUID NOT NULL,

                         is_like bool default false,

                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE swipes ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE swipes ADD FOREIGN KEY (profile_id) REFERENCES user_profiles (id);

-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_swipes_updated_at BEFORE UPDATE ON swipes FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

COMMIT;