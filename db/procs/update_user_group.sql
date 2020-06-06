-- Stored procedure to update a user_group record
 CREATE OR REPLACE FUNCTION users.update_user_group(
	user_ids_in text[],
  group_name_in text
 )
 RETURNS void AS $$
 BEGIN
  DELETE FROM users.user_group;
  INSERT INTO users.user_group (user_id, group_name) (SELECT * FROM unnest(user_ids_in, group_name_in));
 END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;