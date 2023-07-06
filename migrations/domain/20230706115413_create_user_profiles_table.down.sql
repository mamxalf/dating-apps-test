BEGIN;

DROP TABLE IF EXISTS user_profiles CASCADE;
DROP TRIGGER IF EXISTS set_user_profiles_updated_at ON user_profiles CASCADE;

COMMIT;