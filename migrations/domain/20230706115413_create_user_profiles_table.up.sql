BEGIN;

CREATE TABLE IF NOT EXISTS user_profiles (
                                             id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
                                             user_id UUID NOT NULL,

                                             full_name varchar NOT NULL,
                                             age integer NOT NULL,
                                             gender varchar NOT NULL,

                                             created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE user_profiles ADD FOREIGN KEY (user_id) REFERENCES users (id);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_user_profiles_updated_at BEFORE UPDATE ON user_profiles FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

COMMIT;