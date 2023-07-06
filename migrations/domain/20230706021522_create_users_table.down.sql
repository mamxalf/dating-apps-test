BEGIN;

DROP TABLE IF EXISTS users CASCADE;
DROP TRIGGER IF EXISTS set_users_updated_at ON users CASCADE;

DROP TABLE IF EXISTS user_sessions CASCADE;
DROP TRIGGER IF EXISTS set_user_sessions_updated_at ON user_sessions CASCADE;

COMMIT;