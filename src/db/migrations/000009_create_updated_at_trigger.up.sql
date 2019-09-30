CREATE FUNCTION set_update_time() RETURNS opaque AS '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' language 'plpgsql';

CREATE TRIGGER update_tri BEFORE UPDATE ON users FOR EACH row EXECUTE PROCEDURE set_update_time();
CREATE TRIGGER update_tri BEFORE UPDATE ON user_credentials FOR EACH row EXECUTE PROCEDURE set_update_time();
CREATE TRIGGER update_tri BEFORE UPDATE ON user_profiles FOR EACH row EXECUTE PROCEDURE set_update_time();
CREATE TRIGGER update_tri BEFORE UPDATE ON user_access_tokens FOR EACH row EXECUTE PROCEDURE set_update_time();
CREATE TRIGGER update_tri BEFORE UPDATE ON user_refresh_tokens FOR EACH row EXECUTE PROCEDURE set_update_time();
CREATE TRIGGER update_tri BEFORE UPDATE ON user_coordinates FOR EACH row EXECUTE PROCEDURE set_update_time();
