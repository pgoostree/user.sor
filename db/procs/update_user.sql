-- Stored procedure to update a user record
 CREATE OR REPLACE FUNCTION users.update_user(
   user_id_in text,
	first_name_in text,
	last_name_in text
 )
 RETURNS SETOF users.user AS $$
 BEGIN
     RETURN QUERY
     UPDATE users.user SET
        first_name = first_name_in,
        last_name = last_name_in
     WHERE user_id = user_id_in
        RETURNING user_id, first_name, last_name;
 END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;