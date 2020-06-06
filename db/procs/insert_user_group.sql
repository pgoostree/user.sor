-- Stored procedure to insert a user_group record
 CREATE OR REPLACE FUNCTION users.insert_user_group(
	user_id_in text,
  group_name_in text
 )
 RETURNS SETOF users.user_group AS $$
 BEGIN
     RETURN QUERY
     INSERT INTO users.user_group (
       user_id,
       group_name
     )
     VALUES (
       user_id_in,
       group_name_in
     )
        RETURNING user_id, group_name, created, modified;
 END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;