-- Stored procedure to insert a user record
 CREATE OR REPLACE FUNCTION users.insert_user(
	user_id_in text,
	first_name_in text,
	last_name_in text
 )
 RETURNS SETOF users.user AS $$
 BEGIN
     RETURN QUERY
     INSERT INTO users.user (
       user_id,
       first_name,
       last_name
     )
     VALUES (
       user_id_in,
       first_name_in,
       last_name_in
     )
        RETURNING user_id, first_name, last_name;
 END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;